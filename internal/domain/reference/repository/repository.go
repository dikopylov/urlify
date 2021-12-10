package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"urlify/internal/domain/reference/model"
)

const TABLE = "references"

type ReferenceRepository interface {
	Insert(entity *model.Reference)
	GetByHash(hash string) model.Reference
}

type PsqlReferenceRepository struct {
	db *sqlx.DB
}

func (repository PsqlReferenceRepository) Insert(entity *model.Reference) {
	sql := fmt.Sprintf(`INSERT INTO %s (id, url, hash, created_at) VALUES (generateUUIDv4(), :url,:hash, now()) RETURNING id`, TABLE)

	result, err := repository.db.NamedExec(sql, entity)

	if err != nil {
		log.Fatalln(err)
	}

	entity.ID, _ = result.LastInsertId()
}
