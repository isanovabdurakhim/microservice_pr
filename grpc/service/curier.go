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

type curierService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedCurierServiceServer
}

func NewCurierService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *curierService {
	return &curierService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (i *curierService) Create(ctx context.Context, req *pb.Createcurier) (resp *pb.curier, err error) {

	i.log.Info("---Createcurier------>", logger.Any("req", req))

	pKey, err := i.strg.Curier().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!Createcurier->curier->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Curier().GetById(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeycurier->curier->GetByID--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *curierService) GetById(ctx context.Context, req *pb.curierPrimaryKey) (resp *pb.curier, err error) {

	fmt.Println("ok")

	c.log.Info("---GetcurierByID------>", logger.Any("req", req))

	resp, err = c.strg.Curier().GetById(ctx, req)
	if err != nil {
		c.log.Error("!!!GetcurierByID->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *curierService) GetList(ctx context.Context, req *pb.GetListcurierRequest) (resp *pb.GetListcurierResponse, err error) {

	i.log.Info("---GetCategories------>", logger.Any("req", req))

	resp, err = i.strg.Curier().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCategories->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *curierService) Update(ctx context.Context, req *pb.Updatecurier) (resp *pb.curier, err error) {

	i.log.Info("---Updatecurier------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Curier().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!Updatecurier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Curier().GetById(ctx, &pb.curierPrimaryKey{curierId: req.curierId})
	if err != nil {
		i.log.Error("!!!Getcurier->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *curierService) Delete(ctx context.Context, req *pb.curierPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCurier------>", logger.Any("req", req))

	err = i.strg.Curier().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!Deletecurier->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
