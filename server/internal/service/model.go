package service

import (
	"context"

	"github.com/mikeziminio/dynarun/server/internal/domain"
	"go.uber.org/zap"
)

type ModelRepo interface {
	CreateModel(
		ctx context.Context,
		name string,
		repoId string,
		filename string,
		inputTokenPrice int64,
		outputTokenPrice int64,
	) (*domain.Model, error)
	UpdateModel(
		ctx context.Context,
		id domain.ID,
		name *string,
		repoId *string,
		filename *string,
		inputTokenPrice *int64,
		outputTokenPrice *int64,
	) (*domain.Model, error)
	DeleteModel(ctx context.Context, id domain.ID) error
	ListModel(ctx context.Context) ([]domain.Model, error)
}

type ModelService struct {
	logger    *zap.Logger
	modelRepo ModelRepo
}

func NewModelService(
	modelRepo ModelRepo,
	logger *zap.Logger,
) (*ModelService, error) {
	return &ModelService{
		modelRepo: modelRepo,
		logger:    logger,
	}, nil
}

func (s *ModelService) CreateModel(
	ctx context.Context,
	name string,
	repoId string,
	filename string,
	inputTokenPrice int64,
	outputTokenPrice int64,
) (*domain.Model, error) {
	return s.modelRepo.CreateModel(
		ctx,
		name, repoId, filename,
		inputTokenPrice, outputTokenPrice,
	)
}

func (s *ModelService) UpdateModel(
	ctx context.Context,
	id domain.ID,
	name *string,
	repoId *string,
	filename *string,
	inputTokenPrice *int64,
	outputTokenPrice *int64,
) (*domain.Model, error) {
	return s.modelRepo.UpdateModel(
		ctx,
		id, name, repoId, filename,
		inputTokenPrice, outputTokenPrice,
	)
}

func (s *ModelService) DeleteModel(ctx context.Context, id domain.ID) error {
	return s.modelRepo.DeleteModel(ctx, id)
}

func (s *ModelService) ListModel(ctx context.Context) ([]domain.Model, error) {
	return s.modelRepo.ListModel(ctx)
}
