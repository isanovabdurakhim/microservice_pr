package storage

import (
	pb "app/genproto/organization_service"
	"context"
)

type StorageI interface {
	Branch() BranchRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req *pb.CreateBranch) (resp *pb.BranchPrimaryKey, err error)
	GetAll(ctx context.Context, req *pb.GetListBranchRequest) (*pb.GetListBranchResponse, error)
	GetById(ctx context.Context, req *pb.BranchPrimaryKey) (*pb.Branch, error)
	Update(ctx context.Context, req *pb.UpdateBranch) (rowsAffected int64,err error)
	Delete(ctx context.Context, req *pb.BranchPrimaryKey) error
}
