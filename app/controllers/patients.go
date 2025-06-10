package controllers

import (
	"log/slog"
	"net/http"

	"github.com/haleyrc/phir/app/models"
	"github.com/haleyrc/phir/app/routes"
	"github.com/haleyrc/phir/app/services"
	"github.com/haleyrc/phir/lib/web"
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

func (c *PatientsController) New(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := patients.New(models.Patient{}).Render(ctx, w); err != nil {
		c.Logger.ErrorContext(ctx, "patients controller: new", slog.Any("err", err))
	}
}

func NewCreatePatientParams(r *http.Request) services.CreatePatientParams {
	return services.CreatePatientParams{
		FirstName: web.NewString(r.PostFormValue("first_name")),
		LastName:  web.NewString(r.PostFormValue("last_name")),
	}
}

func (c *PatientsController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := NewCreatePatientParams(r)
	result, err := services.NewPatientCreator(c.Repo).CreatePatient(ctx, params)
	if err != nil {
		c.Logger.ErrorContext(ctx, "patients controller: create", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result.Created {
		http.Redirect(w, r, routes.PatientsPath(), http.StatusSeeOther)
		return
	}

	if err := patients.New(result.Patient).Render(ctx, w); err != nil {
		c.Logger.ErrorContext(ctx, "patients controller: new", slog.Any("err", err))
	}
}

func (c *PatientsController) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	patient, err := c.Repo.GetPatient(ctx, web.NewInt(id))
	if err != nil {
		c.Logger.ErrorContext(ctx, "patients controller: show", slog.String("id", id), slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := patients.Show(patient).Render(ctx, w); err != nil {
		c.Logger.ErrorContext(ctx, "patients controller: show", slog.String("id", id), slog.Any("err", err))
	}
}
