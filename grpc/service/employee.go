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

type employeeService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	pb.UnimplementedEmployeeServiceServer
}

func NewEmployeeService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *employeeService {
	return &employeeService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (i *employeeService) Create(ctx context.Context, req *pb.Createemployee) (resp *pb.employee, err error) {

	i.log.Info("---Createemployee------>", logger.Any("req", req))

	pKey, err := i.strg.Employee().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!Createemployee->employee->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Employee().GetById(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyemployee->employee->GetByID--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (c *employeeService) GetById(ctx context.Context, req *pb.employeePrimaryKey) (resp *pb.employee, err error) {

	fmt.Println("ok")

	c.log.Info("---GetemployeeByID------>", logger.Any("req", req))

	resp, err = c.strg.Employee().GetById(ctx, req)
	if err != nil {
		c.log.Error("!!!GetemployeeByID->employee->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *employeeService) GetList(ctx context.Context, req *pb.GetListemployeeRequest) (resp *pb.GetListemployeeResponse, err error) {

	i.log.Info("---GetCategories------>", logger.Any("req", req))

	resp, err = i.strg.Employee().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetCategories->employee->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *employeeService) Update(ctx context.Context, req *pb.Updateemployee) (resp *pb.employee, err error) {

	i.log.Info("---Updateemployee------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Employee().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!Updateemployee--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Employee().GetById(ctx, &pb.employeePrimaryKey{employeeId: req.employeeId})
	if err != nil {
		i.log.Error("!!!Getemployee->employee->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *employeeService) Delete(ctx context.Context, req *pb.employeePrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteEmployee------>", logger.Any("req", req))

	err = i.strg.Employee().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!Deleteemployee->employee->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
