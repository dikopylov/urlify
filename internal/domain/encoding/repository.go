package encoding

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"urlify/internal/domain/model"
)

const TABLE = "references"

type EncodeRepository interface {
	Insert(entity *model.Reference)
	GetByHash(hash string) model.Reference
}

type PsqlEncodeRepository struct {
	db *sqlx.DB
}

func (repository PsqlEncodeRepository) Insert(entity *model.Reference) {
	sql := fmt.Sprintf(`INSERT INTO %s (id, url,hash,created_at) VALUES (generateUUIDv4(), :url,:hash, now()) RETURNING id`, TABLE)

	result, err := repository.db.NamedExec(sql, entity)

	if err != nil {
		log.Fatalln(err)
	}

	entity.ID, _ = result.LastInsertId()
}
