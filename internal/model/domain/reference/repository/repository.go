package repository

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log"
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
	Insert(entity *model.Reference)
	GetByCriteria(criteria Criteria) *model.Reference
}

type PsqlReferenceRepository struct {
	db *sqlx.DB
}

func NewPsqlReferenceRepository(db *sqlx.DB) *PsqlReferenceRepository {
	return &PsqlReferenceRepository{db: db}
}

func (repository *PsqlReferenceRepository) Insert(entity *model.Reference) {
	// @todo
	//query, _, err := sq.Insert(Table).Columns("url", "hash", "created_at").ToSql()

	if err != nil {
		log.Printf("Error: %s\n", err.Error())

		return
	}

	_, err = repository.db.NamedExec(query, entity)

	if err != nil {
		log.Fatalln(err)
	}
}

func (repository *PsqlReferenceRepository) GetByCriteria(criteria Criteria) *model.Reference {
	reference := model.Reference{}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("*").
		From(Table).
		Where(criteria.parameters).
		Limit(1).
		ToSql()

	if err != nil {
		log.Printf("Error: %s\n", err.Error())

		return nil
	}

	err = repository.db.Get(&reference, query, args)

	switch err {
	case nil:
		return &reference
	case sql.ErrNoRows:
		return nil
	default:
		log.Fatalln(err)
		return nil
	}

}
