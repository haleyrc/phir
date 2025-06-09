package controllers

import (
	"log/slog"
	"net/http"

	"github.com/haleyrc/phir/app/models"
	"github.com/haleyrc/phir/templates/pages/patients"
)

type PatientsController struct {
	Logger *slog.Logger
	Repo   models.Repository
}

func NewPatientsController(logger *slog.Logger, repo models.Repository) PatientsController {
	return PatientsController{
		Logger: logger,
		Repo:   repo,
	}
}

func (c *PatientsController) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	allPatients, err := c.Repo.ListPatients(ctx)
	if err != nil {
		c.Logger.ErrorContext(ctx, "patients controller: index", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	props := patients.IndexProps{
		Patients: allPatients,
	}
	if err := patients.Index(props).Render(ctx, w); err != nil {
		c.Logger.ErrorContext(ctx, "patients controller: index", slog.Any("err", err))
	}
}
