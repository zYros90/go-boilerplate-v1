package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/zYros90/go-boilerplate-v1/api/v1/generated"
	"github.com/zYros90/go-boilerplate-v1/app/config"
	"github.com/zYros90/go-boilerplate-v1/app/internal/service"
	"github.com/zYros90/pkg/grpcmw"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	grpcServer *grpc.Server
	listener   net.Listener
	svc        *service.Service
	conf       *config.Config
	logger     *zap.Logger
}

func New(
	conf *config.Config,
	logger *zap.Logger,
	svc *service.Service,
) (*GRPCServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 5000))
	if err != nil {
		return nil, err
	}
	var opts []grpc.ServerOption

	jwtMW := grpcmw.NewJwtMW(conf.Server.JWTSecret, logger)

	opts = append(opts, grpc.UnaryInterceptor(jwtMW))

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	pb.RegisterUserSvcServer(grpcServer, svc)
	// pb.RegisterLoginSvcServer(grpcServer, svc)

	srv := &GRPCServer{
		grpcServer: grpcServer,
		listener:   listener,
		svc:        svc,
		conf:       conf,
		logger:     logger,
	}

	return srv, nil
}

func (srv *GRPCServer) Start(address string) error {
	return srv.grpcServer.Serve(srv.listener)
}

func (srv *GRPCServer) Shutdown(ctx context.Context) error {
	srv.grpcServer.Stop()
	return nil
}
