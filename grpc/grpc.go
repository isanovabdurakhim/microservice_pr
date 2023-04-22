package grpc

import (
	"app/config"
	pb "app/genproto/organization_service"
	"app/grpc/service"
	"app/pkg/logger"
	"app/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	pb.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg))
	pb.RegisterCurierServiceServer(grpcServer, service.NewCurierService(cfg, log, strg))
	pb.RegisterEmployeeServiceServer(grpcServer, service.NewEmployeeService(cfg, log, strg))
	pb.RegisterStoreServiceSErver(grpcServer, service.NewStoreService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
