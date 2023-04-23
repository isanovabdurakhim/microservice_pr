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

type storeService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedStoreServiceServer
}

func NewStoreService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *storeService {
	return &storeService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (i *storeService) Create(ctx context.Context, req *pb.CreateStore) (resp *pb.Store, err error) {

	i.log.Info("---Createstore------>", logger.Any("req", req))

	pKey, err := i.strg.Store().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!Createstore->Store->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Store().GetById(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeystore->Store->GetByID--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *storeService) GetById(ctx context.Context, req *pb.StorePrimaryKey) (resp *pb.Store, err error) {

	fmt.Println("ok")

	c.log.Info("---GetstoreByID------>", logger.Any("req", req))

	resp, err = c.strg.Store().GetById(ctx, req)
	if err != nil {
		c.log.Error("!!!GetstoreByID->Store->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *storeService) GetList(ctx context.Context, req *pb.GetListStoreRequest) (resp *pb.GetListStoreResponse, err error) {

	i.log.Info("---GetCategories------>", logger.Any("req", req))

	resp, err = i.strg.Store().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCategories->Store->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *storeService) Update(ctx context.Context, req *pb.UpdateStore) (resp *pb.Store, err error) {

	i.log.Info("---UpdateStore------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Store().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!Updatestore--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Store().GetById(ctx, &pb.StorePrimaryKey{Id: req.StoreId})
	if err != nil {
		i.log.Error("!!!Getstore->Store->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *storeService) Delete(ctx context.Context, req *pb.StorePrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteStore------>", logger.Any("req", req))

	err = i.strg.Store().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!Deletestore->Store->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
