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

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) storage.StoreRepoI {
	return &storeRepo{
		db: db,
	}
}

func (c *storeRepo) Create(ctx context.Context, req *pb.CreateStore) (resp *pb.StorePrimaryKey, err error) {

	var store_id = uuid.New()

	query := `INSERT INTO "stores" (
				store_id,
				branch_id,
				store_name,
				updated_at
			) VALUES ($1, $2, $3, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		store_id.String(),
		req.BranchId,
		req.StoreName,
	)

	if err != nil {
		return nil, err
	}

	return &pb.StorePrimaryKey{Id: store_id.String()}, nil
}

func (c *storeRepo) GetById(ctx context.Context, req *pb.StorePrimaryKey) (resp *pb.Store, err error) {

	query := `
		SELECT
			store_id,
			branch_id,
			store_name,
			updated_at
			deleted_at
		FROM "stores"
		WHERE store_id = $1 and deleted_at is null
	`

	var (
		store_id    sql.NullString
		branch_id   sql.NullString
		store_name    sql.NullString
		created_at   sql.NullString
		updated_at   sql.NullString
		deleted_at   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&store_id,
		&branch_id,
		&store_name,
		&created_at,
		&updated_at,
		&deleted_at,
	)

	if err != nil {
		return resp, err
	}

	resp = &pb.Store{
		StoreId:  store_id.String,
		BranchId:	branch_id.String,
		StoreName:  store_name.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
		DeletedAt: deleted_at.String,
	}

	return
}

func (c *storeRepo) GetAll(ctx context.Context, req *pb.GetListStoreRequest) (resp *pb.GetListStoreResponse, err error) {

	resp = &pb.GetListStoreResponse{}

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
			store_id,
			branch_id,
			store_name,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS'),
			COALESCE(TO_CHAR(deleted_at, 'YYYY-MM-DD HH24:MI:SS'), '')
		FROM "stores"
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

		var store pb.Store

		err := rows.Scan(
			&resp.Count,
			&store.StoreId,
			&store.BranchId,
			&store.StoreName,
			&store.CreatedAt,
			&store.UpdatedAt,
			&store.DeletedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Stores = append(resp.Stores, &store)
	}

	return
}

func (c *storeRepo) Update(ctx context.Context, req *pb.UpdateStore) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "stores"
			SET
				branch_id = :branch_id,
				store_name = :store_name,
				updated_at = now()
			WHERE
				store_id = :store_id`
	params = map[string]interface{}{
		"store_id":      req.GetStoreId(),
		"branch_id":    req.GetBranchId(),
		"store_name":    req.GetStoreName(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *storeRepo) Delete(ctx context.Context, req *pb.StorePrimaryKey) error {

	query := `DELETE FROM "stores" WHERE store_id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
