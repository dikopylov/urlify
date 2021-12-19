package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"urlify/internal/model/domain/reference/model"
)

const (
	Table      = "reference"
	ColumnUrl  = "url"
	ColumnHash = "hash"
)

type Criteria struct {
	parameters map[string]interface{}
}

func (c *Criteria) AddParameter(column string, value interface{}) {
	if c.parameters == nil {
		c.parameters = make(map[string]interface{})
	}

	c.parameters[column] = value
}

type ReferenceRepository interface {
	Insert(entity *model.Reference) error
	GetByCriteria(criteria Criteria) (*model.Reference, error)
}
type PsqlReferenceRepository struct {
	db *sqlx.DB
}

func NewPsqlReferenceRepository(db *sqlx.DB) *PsqlReferenceRepository {
	return &PsqlReferenceRepository{db: db}
}

func (repository *PsqlReferenceRepository) Insert(entity *model.Reference) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Insert(Table).
		Columns("url", "hash", "created_at").
		Values(entity.Url, entity.Hash, sq.Expr("now()")).
		ToSql()

	if err != nil {
		return err
	}

	_, err = repository.db.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (repository *PsqlReferenceRepository) GetByCriteria(criteria Criteria) (*model.Reference, error) {
	reference := model.Reference{}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("*").
		From(Table).
		Where(criteria.parameters).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = repository.db.Get(&reference, query, args...)

	if err != nil {
		return nil, err
	}

	return &reference, nil
}
