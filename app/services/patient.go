package services

import (
	"context"
	"fmt"

	"github.com/haleyrc/phir/app/models"
)

type PatientCreator struct {
	Repo models.Repository
}

func NewPatientCreator(repo models.Repository) PatientCreator {
	return PatientCreator{
		Repo: repo,
	}
}

type CreatePatientParams struct {
	FirstName string
	LastName  string
}

type CreatePatientResult struct {
	Created bool
	Patient models.Patient
}

func (pc PatientCreator) CreatePatient(ctx context.Context, params CreatePatientParams) (CreatePatientResult, error) {
	patient, err := pc.Repo.CreatePatient(ctx, models.CreatePatientParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
	})
	if err != nil {
		return CreatePatientResult{}, fmt.Errorf("patient service: create patient: %w", err)
	}

	if !patient.Valid() {
		return CreatePatientResult{Created: false, Patient: patient}, nil
	}

	return CreatePatientResult{Created: true, Patient: patient}, nil
}
