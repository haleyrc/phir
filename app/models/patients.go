package models

import (
	"context"
	"fmt"
	"strings"

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

type CreatePatientParams struct {
	FirstName string
	LastName  string
}

func (repo *Repository) CreatePatient(ctx context.Context, params CreatePatientParams) (Patient, error) {
	params.FirstName = strings.TrimSpace(params.FirstName)
	params.LastName = strings.TrimSpace(params.LastName)

	var patient Patient
	if len(params.FirstName) == 0 {
		patient.Errors.FirstName.Append("Must be one or more characters.")
	}
	if len(params.LastName) == 0 {
		patient.Errors.LastName.Append("Must be one or more characters.")
	}

	if !patient.Valid() {
		return patient, nil
	}

	var err error
	patient.Patient, err = repo.q.CreatePatient(ctx, db.CreatePatientParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
	})
	if err != nil {
		return Patient{}, fmt.Errorf("db: create patient: %w", err)
	}

	return patient, nil
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
