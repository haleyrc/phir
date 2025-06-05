package controllers

import (
	"log/slog"
	"net/http"

	"github.com/haleyrc/phir/templates/pages"
)

type AppController struct {
	Logger *slog.Logger
}

func NewAppController(logger *slog.Logger) AppController {
	return AppController{
		Logger: logger,
	}
}

func (c *AppController) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	props := pages.IndexProps{}
	if err := pages.Index(props).Render(ctx, w); err != nil {
		c.Logger.ErrorContext(ctx, "app controller: index", slog.Any("err", err))
	}
}
