package postgres

import (
	"app/pkg/helper"
	"app/storage"
	"context"
	"database/sql"

	pb "app/genproto/organization_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type courierRepo struct {
	db *pgxpool.Pool
}

func NewCourierRepo(db *pgxpool.Pool) storage.CourierRepoI {
	return &courierRepo{
		db: db,
	}
}

func (c *courierRepo) Create(ctx context.Context, req *pb.CreateCourier) (resp *pb.CourierPrimaryKey, err error) {

	var courier_id = uuid.New()

	query := `INSERT INTO "couriers" (
				courier_id,
				first_name,
				last_name,
				phone_number,
				updated_at
			) VALUES ($1, $2, $3, $4, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		courier_id.String(),
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &pb.CourierPrimaryKey{Id: courier_id.String()}, nil
}

func (c *courierRepo) GetById(ctx context.Context, req *pb.CourierPrimaryKey) (resp *pb.Courier, err error) {

	query := `
		SELECT
			curier_id,
			first_name,
			last_name,
			phone_number,
			updated_at
			deleted_at
		FROM "curiers"
		WHERE curier_id = $1 and deleted_at is null
	`

	var (
		courier_id    sql.NullString
		first_name   sql.NullString
		last_name    sql.NullString
		phone_number sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
		deleted_at   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&courier_id,
		&first_name,
		&last_name,
		&phone_number,
		&created_at,
		&updated_at,
		&deleted_at,
	)

	if err != nil {
		return resp, err
	}

	resp = &pb.Courier{
		CourierId:  courier_id.String,
		FirstName: first_name.String,
		LastName:  last_name.String,
		PhoneNumber: phone_number.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
		DeletedAt: deleted_at.String,
	}

	return
}

func (c *courierRepo) GetAll(ctx context.Context, req *pb.GetListCourierRequest) (resp *pb.GetListCourierResponse, err error) {

	resp = &pb.GetListCourierResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " AND TRUE"
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			curier_id,
			first_name,
			last_name,
			phone_number,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'),
			COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')
		FROM "curiers"
		WHERE COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')=''
	`

	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}

	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}

	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {

		var courier pb.Courier

		err := rows.Scan(
			&resp.Count,
			&courier.CourierId,
			&courier.FirstName,
			&courier.LastName,
			&courier.PhoneNumber,
			&courier.CreatedAt,
			&courier.UpdatedAt,
			&courier.DeletedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Couriers = append(resp.Couriers, &courier)
	}

	return
}

func (c *courierRepo) Update(ctx context.Context, req *pb.UpdateCourier) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "curiers"
			SET
				first_name = :first_name,
				last_name = :last_name,
				phone_number = :phone_number,
				updated_at = now()
			WHERE
				curier_id = :curier_id`
	params = map[string]interface{}{
		"courier_id":      req.GetCourierId(),
		"first_name":    req.GetFirstName(),
		"last_name":    req.GetLastName(),
		"phone_number": req.GetPhoneNumber(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *courierRepo) Delete(ctx context.Context, req *pb.CourierPrimaryKey) error {

	query := `DELETE FROM "curiers" WHERE curier_id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
