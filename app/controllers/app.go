package controllers

import (
	"log/slog"
	"net/http"
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
