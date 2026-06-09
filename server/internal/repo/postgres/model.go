package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/mikeziminio/dynarun/server/internal/domain"
)

type ModelRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewModelRepo(pool *pgxpool.Pool, logger *zap.Logger) (*ModelRepo, error) {
	return &ModelRepo{
		pool:   pool,
		logger: logger,
	}, nil
}

func (r *ModelRepo) CreateModel(
	ctx context.Context,
	name string,
	repoID string,
	filename string,
	inputTokenPrice int64,
	outputTokenPrice int64,
) (*domain.Model, error) {
	const q = `
        INSERT INTO model (name, repo_id, filename, input_token_price, output_token_price, updated_at)
        VALUES (@name, @repo_id, @filename, @input_token_price, @output_token_price, now())
        RETURNING id, name, repo_id, filename, input_token_price, output_token_price, updated_at
    `
	r.logger.Debug("query", zap.String("raw", q))
	args := pgx.NamedArgs{
		"name":               name,
		"repo_id":            repoID,
		"filename":           filename,
		"input_token_price":  inputTokenPrice,
		"output_token_price": outputTokenPrice,
	}
	var result domain.Model
	err := r.pool.QueryRow(ctx, q, args).Scan(
		&result.ID,
		&result.Name,
		&result.RepoID,
		&result.Filename,
		&result.InputTokenPrice,
		&result.OutputTokenPrice,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create model: %w", err)
	}

	return &result, nil
}

func (r *ModelRepo) UpdateModel(
	ctx context.Context,
	id domain.ID,
	name *string,
	repoID *string,
	filename *string,
	inputTokenPrice *int64,
	outputTokenPrice *int64,
) (*domain.Model, error) {
	var qb strings.Builder

	args := pgx.NamedArgs{
		"id": id.String(),
	}
	qb.WriteString(`
		UPDATE model
		SET updated_at = now()
	`)
	if name != nil {
		args["name"] = *name
		qb.WriteString(`, name = @name`)
	}
	if repoID != nil {
		args["repo_id"] = *repoID
		qb.WriteString(`, repo_id = @repo_id`)
	}
	if filename != nil {
		args["filename"] = *filename
		qb.WriteString(`, filename = @filename`)
	}
	if inputTokenPrice != nil {
		args["input_token_price"] = *inputTokenPrice
		qb.WriteString(`, input_token_price = @input_token_price`)
	}
	if outputTokenPrice != nil {
		args["output_token_price"] = *outputTokenPrice
		qb.WriteString(`, output_token_price = @output_token_price`)
	}
	qb.WriteString(`
		WHERE id = @id
		RETURNING id, name, repo_id, filename, input_token_price, output_token_price, updated_at
	`)
	q := qb.String()
	r.logger.Debug("query", zap.String("raw", q))

	var model domain.Model
	err := r.pool.QueryRow(ctx, q, args).Scan(
		&model.ID,
		&model.Name,
		&model.RepoID,
		&model.Filename,
		&model.InputTokenPrice,
		&model.OutputTokenPrice,
		&model.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update model: %w", err)
	}
	return &model, nil
}

func (r *ModelRepo) DeleteModel(ctx context.Context, id domain.ID) error {
	const q = `
        DELETE FROM model WHERE id = @id
    `
	r.logger.Debug("query", zap.String("raw", q))
	args := pgx.NamedArgs{
		"id": id,
	}
	tag, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("failed to delete model: %w", err)
	}
	ra := tag.RowsAffected()
	if ra != 1 {
		return fmt.Errorf("affect %d rows instead of 1", ra)
	}
	return nil
}

func (r *ModelRepo) ListModel(ctx context.Context) ([]domain.Model, error) {
	const q = `
        SELECT id, name, repo_id, filename, input_token_price, output_token_price, updated_at
        FROM model
    `
	r.logger.Debug("query", zap.String("raw", q))
	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %w", err)
	}
	models, err := pgx.CollectRows[domain.Model](rows, func(row pgx.CollectableRow) (domain.Model, error) {
		var model domain.Model
		err := row.Scan(
			&model.ID,
			&model.Name,
			&model.RepoID,
			&model.Filename,
			&model.InputTokenPrice,
			&model.OutputTokenPrice,
			&model.UpdatedAt,
		)
		return model, err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list models: %w", err)
	}
	return models, nil
}
