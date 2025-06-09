package models

import (
	"github.com/haleyrc/phir/app/db"
	"github.com/haleyrc/phir/lib/sqlite"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(conn *sqlite.Conn) Repository {
	repo := Repository{
		q: db.New(conn),
	}
	return repo
}

type Errors []string

func (errs *Errors) Append(e string) {
	*errs = append(*errs, e)
}
