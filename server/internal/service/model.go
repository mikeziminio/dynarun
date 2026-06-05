package service

import (
	"github.com/mikeziminio/dynarun/server/internal/domain"
	"go.uber.org/zap"
)

type ModelRepo interface {
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
	name string,
	repoId string,
	filename string,
	inputTokenPrice int64,
	outputTokenPrice int64,
) (*domain.Model, error) {
	return s.modelRepo.CreateModel(
		name, repoId, filename,
		inputTokenPrice, outputTokenPrice,
	)
}

func (s *ModelService) UpdateModel(
	id domain.ID,
	name *string,
	repoId *string,
	filename *string,
	inputTokenPrice *int64,
	outputTokenPrice *int64,
) (*domain.Model, error) {
	return s.modelRepo.UpdateModel(
		id, name, repoId, filename,
		inputTokenPrice, outputTokenPrice,
	)
}

func (s *ModelService) DeleteModel(id domain.ID) error {
	return s.modelRepo.DeleteModel(id)
}

func (s *ModelService) ListModel() ([]domain.Model, error) {
	return s.modelRepo.ListModel()
}
