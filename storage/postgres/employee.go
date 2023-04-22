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

type employeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) storage.EmployeeRepoI {
	return &employeeRepo{
		db: db,
	}
}

func (c *employeeRepo) Create(ctx context.Context, req *pb.CreateEmployee) (resp *pb.EmployeePrimaryKey, err error) {

	var employee_id = uuid.New()

	query := `INSERT INTO "employees" (
				employee_id,
				store_id,
				first_name,
				last_name,
				phone_number,
				updated_at
			) VALUES ($1, $2, $3, $4, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		employee_id.String(),
		req.StoreId,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &pb.EmployeePrimaryKey{EmployeeId: employee_id.String()}, nil
}

func (c *employeeRepo) GetById(ctx context.Context, req *pb.EmployeePrimaryKey) (resp *pb.Employee, err error) {

	query := `
		SELECT
			employee_id,
			store_id,
			first_name,
			last_name,
			phone_number,
			updated_at
			deleted_at
		FROM "employees"
		WHERE employee_id = $1 and deleted_at is null
	`

	var (
		employee_id  sql.NullString
		store_id     sql.NullString
		first_name   sql.NullString
		last_name    sql.NullString
		phone_number sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
		deleted_at   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&employee_id,
		&store_id,
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

	resp = &pb.Employee{
		EmployeeId:  employee_id.String,
		StoreId:     store_id.String,
		FirstName:   first_name.String,
		LastName:    last_name.String,
		PhoneNumber: phone_number.String,
		CreatedAt:   created_at.String,
		UpdatedAt:   updated_at.String,
		DeletedAt:   deleted_at.String,
	}

	return
}

func (c *employeeRepo) GetAll(ctx context.Context, req *pb.GetListEmployeeRequest) (resp *pb.GetListEmployeeResponse, err error) {

	resp = &pb.GetListEmployeeResponse{}

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
			employee_id,
			store_id,
			first_name,
			last_name,
			phone_number,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'),
			COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')
		FROM "employees"
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

		var employee pb.Employee

		err := rows.Scan(
			&resp.Count,
			&employee.EmployeeId,
			&employee.StoreId,
			&employee.FirstName,
			&employee.LastName,
			&employee.PhoneNumber,
			&employee.CreatedAt,
			&employee.UpdatedAt,
			&employee.DeletedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Employees = append(resp.Employees, &employee)
	}

	return
}

func (c *employeeRepo) Update(ctx context.Context, req *pb.UpdateEmployee) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "employees"
			SET
				store_id = :store_id,
				first_name = :first_name,
				last_name = :last_name,
				phone_number = :phone_number,
				updated_at = now()
			WHERE
				employee_id = :employee_id`
	params = map[string]interface{}{
		"employee_id":  req.GetEmployeeId(),
		"store_id":     req.GetStoreId(),
		"first_name":   req.GetFirstName(),
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

func (c *employeeRepo) Delete(ctx context.Context, req *pb.EmployeePrimaryKey) error {

	query := `DELETE FROM "employees" WHERE employee_id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
