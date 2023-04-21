package service

import (
	"app/config"
	pb "app/genproto/organization_service"
	"app/pkg/logger"
	"app/storage"
	"context"
)

type professionService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *professionService {
	return &professionService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *professionService) Create(ctx context.Context, req *pb.CreateBranch) (*pb.Branch, error) {
	branch_id, err := s.strg.Branch().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateBranch", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &pb.Branch{
		BranchId:      branch_id.Id,
		BranchCode:    req.BranchCode,
		BranchName:    req.BranchName,
		BranchAddress: req.BranchAddress,
	}, nil
}

func (s *professionService) GetAll(ctx context.Context, req *pb.GetListBranchRequest) (*pb.GetListBranchResponse, error) {
	resp, err := s.strg.Branch().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllBranch", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}
