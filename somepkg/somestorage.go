package somepkg

import (
	"demomockery/somepkg/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

const queryGetData = `select id, val from sometable where key1 = $1 order by id`

func (s *Storage) GetData(key int) ([]models.DataRec, error) {
	var rec []models.DataRec
	err := s.db.Select(&rec, queryGetData, key)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read record from db")
	}
	return rec, nil
}
