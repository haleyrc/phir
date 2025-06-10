package main

import (
	"context"
	_ "embed"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/haleyrc/phir/app/controllers"
	"github.com/haleyrc/phir/app/models"
	"github.com/haleyrc/phir/lib/sqlite"
	"github.com/haleyrc/server"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	conn, err := sqlite.Open("dev.db")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := models.NewRepository(conn)

	appController := controllers.NewAppController(logger, repo)
	patientsController := controllers.NewPatientsController(logger, repo)

	router := http.NewServeMux()
	router.HandleFunc("GET /patients/new", patientsController.New)
	router.HandleFunc("GET /patients/{id}", patientsController.Show)
	router.HandleFunc("GET /patients", patientsController.Index)
	router.HandleFunc("POST /patients", patientsController.Create)
	router.HandleFunc("GET /{$}", appController.Index)

	server := server.New(os.Getenv("PORT"), router)
	logger.Info("server listening", slog.String("addr", server.Addr()))
	if err := server.ListenAndServe(ctx); err != nil {
		logger.Error("server quit unexpectedly", slog.Any("error", err))
	}
}
