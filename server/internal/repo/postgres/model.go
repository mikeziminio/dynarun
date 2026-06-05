package postgres

import (
	"github.com/mikeziminio/dynarun/server/internal/domain"
	"go.uber.org/zap"
)

type ModelRepo struct {
	logger *zap.Logger
}

func NewModelRepo(logger *zap.Logger) (*ModelRepo, error) {
	return &ModelRepo{
		logger: logger,
	}, nil
}

func (r *ModelRepo) CreateModel(
	name string,
	repoId string,
	filename string,
	inputTokenPrice int64,
	outputTokenPrice int64,
) (*domain.Model, error) {
	return nil, nil
}

func (r *ModelRepo) UpdateModel(
	id domain.ID,
	name *string,
	repoId *string,
	filename *string,
	inputTokenPrice *int64,
	outputTokenPrice *int64,
) (*domain.Model, error) {
	return nil, nil
}

func (r *ModelRepo) DeleteModel(id domain.ID) error {
	return nil
}

func (r *ModelRepo) ListModel() ([]domain.Model, error) {
	return nil, nil
}
