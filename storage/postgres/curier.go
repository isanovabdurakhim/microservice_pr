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

type curierRepo struct {
	db *pgxpool.Pool
}

func NewCurierRepo(db *pgxpool.Pool) storage.CurierRepoI {
	return &curierRepo{
		db: db,
	}
}

func (c *curierRepo) Create(ctx context.Context, req *pb.CreateCurier) (resp *pb.CurierPrimaryKey, err error) {

	var curier_id = uuid.New()

	query := `INSERT INTO "curiers" (
				curier_id,
				first_name,
				last_name,
				phone_number,
				updated_at
			) VALUES ($1, $2, $3, $4, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		curier_id.String(),
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &pb.CurierPrimaryKey{CurierId: curier_id.String()}, nil
}

func (c *curierRepo) GetById(ctx context.Context, req *pb.CurierPrimaryKey) (resp *pb.Curier, err error) {

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
		curier_id    sql.NullString
		first_name   sql.NullString
		last_name    sql.NullString
		phone_number sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
		deleted_at   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&curier_id,
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

	resp = &pb.Curier{
		CurierId:  curier_id.String,
		FirstName: first_name.String,
		LastName:  last_name.String,
		PhoneNumber: phone_number.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
		DeletedAt: deleted_at.String,
	}

	return
}

func (c *curierRepo) GetAll(ctx context.Context, req *pb.GetListCurierRequest) (resp *pb.GetListCurierResponse, err error) {

	resp = &pb.GetListCurierResponse{}

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

		var curier pb.Curier

		err := rows.Scan(
			&resp.Count,
			&curier.CurierId,
			&curier.FirstName,
			&curier.LastName,
			&curier.PhoneNumber,
			&curier.CreatedAt,
			&curier.UpdatedAt,
			&curier.DeletedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Curiers = append(resp.Curiers, &curier)
	}

	return
}

func (c *curierRepo) Update(ctx context.Context, req *pb.UpdateCurier) (rowsAffected int64, err error) {

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
		"curier_id":      req.GetCurierId(),
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

func (c *curierRepo) Delete(ctx context.Context, req *pb.CurierPrimaryKey) error {

	query := `DELETE FROM "curiers" WHERE curier_id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
