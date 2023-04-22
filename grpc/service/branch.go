package service

import (
	"app/config"
	pb "app/genproto/organization_service"
	"app/pkg/logger"
	"app/storage"
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type branchService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *branchService {
	return &branchService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (i *branchService) Create(ctx context.Context, req *pb.CreateBranch) (resp *pb.Branch, err error) {

	i.log.Info("---CreateBranch------>", logger.Any("req", req))

	pKey, err := i.strg.Branch().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateBranch->Branch->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Branch().GetById(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByIdBranch->Branch->GetByID--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *branchService) GetById(ctx context.Context, req *pb.BranchPrimaryKey) (resp *pb.Branch, err error) {

	fmt.Println("salom")

	c.log.Info("---GetBranchByID------>", logger.Any("req", req))

	resp, err = c.strg.Branch().GetById(ctx, req)
	if err != nil {
		c.log.Error("!!!GetBranchByID->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *branchService) GetList(ctx context.Context, req *pb.GetListBranchRequest) (resp *pb.GetListBranchResponse, err error) {

	i.log.Info("---GetCategories------>", logger.Any("req", req))

	resp, err = i.strg.Branch().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCategories->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *branchService) Update(ctx context.Context, req *pb.UpdateBranch) (resp *pb.Branch, err error) {

	i.log.Info("---UpdateBranch------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Branch().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Branch().GetById(ctx, &pb.BranchPrimaryKey{BranchId: req.BranchId})
	if err != nil {
		i.log.Error("!!!GetBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *branchService) Delete(ctx context.Context, req *pb.BranchPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteBranch------>", logger.Any("req", req))

	err = i.strg.Branch().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
