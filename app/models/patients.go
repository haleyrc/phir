package models

import (
	"context"
	"fmt"

	"github.com/haleyrc/phir/app/db"
)

type Patient struct {
	db.Patient

	Errors struct {
		FirstName Errors
		LastName  Errors
	}
}

func (p Patient) Valid() bool {
	return len(p.Errors.FirstName) == 0 && len(p.Errors.LastName) == 0
}

func (repo *Repository) ListPatients(ctx context.Context) ([]Patient, error) {
	ws, err := repo.q.ListPatients(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo: list patients: %w", err)
	}

	patients := make([]Patient, 0, len(ws))
	for _, w := range ws {
		patients = append(patients, Patient{Patient: w})
	}

	return patients, nil
}
