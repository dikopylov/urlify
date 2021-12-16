package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"urlify/internal/domain/reference/model"
)

const (
	TABLE       = "references"
	COLUMN_URL  = "url"
	COLUMN_HASH = "hash"
)

type Criteria struct {
	parameters map[string]interface{}
}

func (c Criteria) AddParameter(column string, value interface{}) {
	c.parameters[column] = value
}

func (c Criteria) convertToWhereClause() string {

	result := ""

	for column, value := range c.parameters {
		result += fmt.Sprintf(`%s=%s`, column, value)
	}

	if result != "" {
		return " WHERE " + result
	}

	return result
}

type ReferenceRepository interface {
	Insert(entity *model.Reference)
	GetByCriteria(criteria Criteria) *model.Reference
}

type PsqlReferenceRepository struct {
	db *sqlx.DB
}

func NewPsqlReferenceRepository(db *sqlx.DB) PsqlReferenceRepository {
	return PsqlReferenceRepository{db: db}
}

func (repository PsqlReferenceRepository) Insert(entity *model.Reference) {
	query := fmt.Sprintf(`INSERT INTO %s (url, hash, created_at) VALUES (:url, :hash, now())`, TABLE)

	_, err := repository.db.NamedExec(query, entity)

	if err != nil {
		log.Fatalln(err)
	}
}

func (repository PsqlReferenceRepository) GetByCriteria(criteria Criteria) *model.Reference {
	reference := model.Reference{}

	query := fmt.Sprintf(`SELECT * FROM %s`, TABLE) + criteria.convertToWhereClause()

	err := repository.db.Get(&reference, query)

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
