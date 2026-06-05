package api

import (
	"context"
	"fmt"

	"github.com/mikeziminio/dynarun/server/internal/domain"
	"github.com/mikeziminio/dynarun/shared/proto/model"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ModelService interface {
	CreateModel(
		name string,
		repoId string,
		filename string,
		inputTokenPrice int64,
		outputTokenPrice int64,
	) (*domain.Model, error)
	UpdateModel(
		id domain.ID,
		name *string,
		repoId *string,
		filename *string,
		inputTokenPrice *int64,
		outputTokenPrice *int64,
	) (*domain.Model, error)
	DeleteModel(id domain.ID) error
	ListModel() ([]domain.Model, error)
}

type ModelServer struct {
	model.UnimplementedModelServiceServer
	modelService ModelService
	logger       *zap.Logger
}

func NewModelServer(
	modelService ModelService,
	logger *zap.Logger,
) *ModelServer {
	return &ModelServer{
		modelService: modelService,
		logger:       logger,
	}
}

func (s *ModelServer) CreateModel(ctx context.Context, req *model.CreateModelRequest) (*model.ModelInfo, error) {
	m, err := s.modelService.CreateModel(
		req.Name,
		req.RepoId,
		req.Filename,
		req.InputTokenPrice,
		req.OutputTokenPrice,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}
	return &model.ModelInfo{
		Id:               m.Id.String(),
		Name:             m.Name,
		RepoId:           m.RepoId,
		Filename:         m.Filename,
		InputTokenPrice:  m.InputTokenPrice,
		OutputTokenPrice: m.OutputTokenPrice,
		UpdatedAt:        timestamppb.New(m.UpdatedAt),
	}, nil
}

func (s *ModelServer) UpdateModel(ctx context.Context, req *model.UpdateModelRequest) (*model.ModelInfo, error) {
	id, err := domain.ParseID(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to parse id: %v", err))
	}
	m, err := s.modelService.UpdateModel(
		id,
		req.Name,
		req.RepoId,
		req.Filename,
		req.InputTokenPrice,
		req.OutputTokenPrice,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}
	return &model.ModelInfo{
		Id:               m.Id.String(),
		Name:             m.Name,
		RepoId:           m.RepoId,
		Filename:         m.Filename,
		InputTokenPrice:  m.InputTokenPrice,
		OutputTokenPrice: m.OutputTokenPrice,
		UpdatedAt:        timestamppb.New(m.UpdatedAt),
	}, nil
}

func (s *ModelServer) DeleteModel(ctx context.Context, req *model.DeleteModelRequest) (*emptypb.Empty, error) {
	id, err := domain.ParseID(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to parse id: %v", err))
	}
	err = s.modelService.DeleteModel(id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}
	return &emptypb.Empty{}, nil
}
func (s *ModelServer) ListModel(ctx context.Context, req *emptypb.Empty) (*model.ModelItems, error) {
	ms, err := s.modelService.ListModel()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}
	modelInfos := make([]*model.ModelInfo, 0)
	for _, m := range ms {
		modelInfos = append(modelInfos, &model.ModelInfo{
			Id:               m.Id.String(),
			Name:             m.Name,
			RepoId:           m.RepoId,
			Filename:         m.Filename,
			InputTokenPrice:  m.InputTokenPrice,
			OutputTokenPrice: m.OutputTokenPrice,
			UpdatedAt:        timestamppb.New(m.UpdatedAt),
		})
	}
	return &model.ModelItems{
		Models: modelInfos,
	}, nil
}
