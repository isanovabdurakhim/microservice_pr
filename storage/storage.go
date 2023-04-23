package storage

import (
	pb "app/genproto/organization_service"
	"context"
)

type StorageI interface {
	Branch() BranchRepoI
	Employee() EmployeeRepoI
	Courier() CourierRepoI
	Store() StoreRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req *pb.CreateBranch) (resp *pb.BranchPrimaryKey, err error)
	GetAll(ctx context.Context, req *pb.GetListBranchRequest) (*pb.GetListBranchResponse, error)
	GetById(ctx context.Context, req *pb.BranchPrimaryKey) (*pb.Branch, error)
	Update(ctx context.Context, req *pb.UpdateBranch) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.BranchPrimaryKey) error
}

type CourierRepoI interface {
	Create(ctx context.Context, req *pb.CreateCourier) (resp *pb.CourierPrimaryKey, err error)
	GetAll(ctx context.Context, req *pb.GetListCourierRequest) (*pb.GetListCourierResponse, error)
	GetById(ctx context.Context, req *pb.CourierPrimaryKey) (*pb.Courier, error)
	Update(ctx context.Context, req *pb.UpdateCourier) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.CourierPrimaryKey) error
}

type StoreRepoI interface {
	Create(ctx context.Context, req *pb.CreateStore) (resp *pb.StorePrimaryKey, err error)
	GetAll(ctx context.Context, req *pb.GetListStoreRequest) (*pb.GetListStoreResponse, error)
	GetById(ctx context.Context, req *pb.StorePrimaryKey) (*pb.Store, error)
	Update(ctx context.Context, req *pb.UpdateStore) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.StorePrimaryKey) error
}

type EmployeeRepoI interface {
	Create(ctx context.Context, req *pb.CreateEmployee) (resp *pb.EmployeePrimaryKey, err error)
	GetAll(ctx context.Context, req *pb.GetListEmployeeRequest) (*pb.GetListEmployeeResponse, error)
	GetById(ctx context.Context, req *pb.EmployeePrimaryKey) (*pb.Employee, error)
	Update(ctx context.Context, req *pb.UpdateEmployee) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *pb.EmployeePrimaryKey) error
}



