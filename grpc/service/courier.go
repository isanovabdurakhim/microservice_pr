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

type courierService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedCourierServiceServer
}

func NewCourierService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *courierService {
	return &courierService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (i *courierService) Create(ctx context.Context, req *pb.CreateCourier) (resp *pb.Courier, err error) {

	i.log.Info("---Createcurier------>", logger.Any("req", req))

	pKey, err := i.strg.Courier().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!Createcurier->curier->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Courier().GetById(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeycurier->curier->GetByID--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *courierService) GetById(ctx context.Context, req *pb.CourierPrimaryKey) (resp *pb.Courier, err error) {

	fmt.Println("ok")

	c.log.Info("---GetcurierByID------>", logger.Any("req", req))

	resp, err = c.strg.Courier().GetById(ctx, req)
	if err != nil {
		c.log.Error("!!!GetcurierByID->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *courierService) GetList(ctx context.Context, req *pb.GetListCourierRequest) (resp *pb.GetListCourierResponse, err error) {

	i.log.Info("---GetCategories------>", logger.Any("req", req))

	resp, err = i.strg.Courier().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCategories->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *courierService) Update(ctx context.Context, req *pb.UpdateCourier) (resp *pb.Courier, err error) {

	i.log.Info("---Updatecurier------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Courier().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!Updatecurier--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Courier().GetById(ctx, &pb.CourierPrimaryKey{Id: req.CourierId})
	if err != nil {
		i.log.Error("!!!Getcurier->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *courierService) Delete(ctx context.Context, req *pb.CourierPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteCourier------>", logger.Any("req", req))

	err = i.strg.Courier().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!Deletecurier->curier->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
