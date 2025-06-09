package controllers

import (
	"log/slog"
	"net/http"

	"github.com/haleyrc/phir/app/models"
	"github.com/haleyrc/phir/templates/pages"
)

type AppController struct {
	Logger *slog.Logger
	Repo   models.Repository
}

func NewAppController(logger *slog.Logger, repo models.Repository) AppController {
	return AppController{
		Logger: logger,
		Repo:   repo,
	}
}

func (c *AppController) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	allPatients, err := c.Repo.ListPatients(ctx)
	if err != nil {
		c.Logger.ErrorContext(ctx, "app controller: index", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	props := pages.IndexProps{
		Patients: allPatients,
	}
	if err := pages.Index(props).Render(ctx, w); err != nil {
		c.Logger.ErrorContext(ctx, "app controller: index", slog.Any("err", err))
	}
}
