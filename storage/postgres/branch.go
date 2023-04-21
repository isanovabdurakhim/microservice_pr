package postgres

import (
	"app/pkg/helper"
	"app/storage"
	"context"
	"database/sql"
	"fmt"

	pb "app/genproto/organization_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepoI {
	return &branchRepo{
		db: db,
	}
}

func (r *branchRepo) Create(ctx context.Context, req *pb.CreateBranch) (resp *pb.BranchPrimaryKey, err error) {

	var branch_id = uuid.New()

	query := `
		INSERT INTO "branches" (
			branch_id,
			branch_code,
			branch_name,
			branch_address,
			updated_at
		) 
		 VALUES ($1, $2, $3, $4, now())
	`

	_, err = r.db.Exec(
		ctx,
		query,
		branch_id.String(),
		req.BranchCode,
		req.BranchName,
		req.BranchAddress,
	)

	if err != nil {
		return nil, err
	}

	return &pb.BranchPrimaryKey{Id: branch_id.String()}, nil
}

func (r *branchRepo) GetById(ctx context.Context, req *pb.BranchPrimaryKey) (resp *pb.Branch, err error) {
	query := `
		SELECT 	
			branch_id,
			branch_code,
			branch_name,
			branch_address,
			created_at,
			updated_at
		FROM "branches"
		WHERE branch_id $1
	`

	var (
		branch_id      sql.NullString
		branch_code    sql.NullString
		branch_name    sql.NullString
		branch_address sql.NullString
		createdAt      sql.NullString
		updatedAt      sql.NullString
	)

	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&branch_id,
		&branch_code,
		&branch_name,
		&branch_address,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &pb.Branch{
		BranchId:      branch_id.String,
		BranchCode:    branch_code.String,
		BranchName:    branch_name.String,
		BranchAddress: branch_address.String,
		CreatedAt:     createdAt.String,
		UpdatedAt:     updatedAt.String,
	}

	return
}

func (r *branchRepo) GetAll(ctx context.Context, req *pb.GetListBranchRequest) (*pb.GetListBranchResponse, error) {
	var (
		resp   pb.GetListBranchResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM branch WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				name
			FROM branch
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var (
			branch_id      sql.NullString
			branch_code    sql.NullString
			branch_name    sql.NullString
			branch_address sql.NullString
			createdAt      sql.NullString
			updatedAt      sql.NullString
		)

		err = rows.Scan(
			&branch_id,
			&branch_code,
			&branch_name,
			&branch_address,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return &resp, err
		}

		resp.Branches = append(resp.Branches, &pb.Branch{
			BranchId:      branch_id.String,
			BranchCode:    branch_code.String,
			BranchName:    branch_name.String,
			BranchAddress: branch_address.String,
			CreatedAt:     createdAt.String,
			UpdatedAt:     updatedAt.String,
		})
	}

	return &resp, nil
}

func (r *branchRepo) Update(ctx context.Context, req *pb.UpdateBranch) (rowsAffected int64,err error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "branches"
			SET
				branch_code = :branch_code,
				branch_name = :branch_name,
				branch_address = :branch_address,
				updated_at = now()
			WHERE
				branch_id = :branch_id`
	params = map[string]interface{}{
		"branch_id":      req.GetBranchId(),
		"branch_code":    req.GetBranchCode(),
		"branch_name":    req.GetBranchName(),
		"branch_address": req.GetBranchAddress(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (r *branchRepo) Delete(ctx context.Context, req *pb.BranchPrimaryKey) error {
	query := `DELETE FROM "branches" WHERE branch_id = $1`

	_, err := r.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}