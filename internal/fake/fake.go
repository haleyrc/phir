package fake

import (
	"context"
	"fmt"

	"github.com/haleyrc/phir/app/db"
	"github.com/haleyrc/phir/lib/sqlite"
)

type Faker struct {
	q *db.Queries
}

func New(conn *sqlite.Conn) Faker {
	return Faker{q: db.New(conn)}
}

type Patient struct {
	FirstName string
	LastName  string
}

func (f Faker) CreatePatient(ctx context.Context, p Patient) int64 {
	patient, err := f.q.CreatePatient(ctx, db.CreatePatientParams{
		FirstName: p.FirstName,
		LastName:  p.LastName,
	})
	if err != nil {
		fail("create patient", err)
	}

	return patient.ID
}

func fail(label string, err error) {
	err = fmt.Errorf("faker: %s: %v", label, err)
	panic(err)
}
