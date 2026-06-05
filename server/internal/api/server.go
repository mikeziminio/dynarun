package api

import (
	"net"

	"github.com/mikeziminio/dynarun/shared/proto/model"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	address      string
	modelService ModelService
	logger       *zap.Logger
}

func NewServer(
	address string,
	modelService ModelService,
	logger *zap.Logger,
) *Server {
	return &Server{
		address:      address,
		modelService: modelService,
		logger:       logger,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	modelServer := NewModelServer(s.modelService, s.logger)
	model.RegisterModelServiceServer(srv, modelServer)
	reflection.Register(srv)
	s.logger.Info("Start to serve grpc", zap.String("address", s.address))
	return srv.Serve(lis)
}
