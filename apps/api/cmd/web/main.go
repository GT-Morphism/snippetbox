package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"snippetbox.gentiluomo.dev/internal/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	connStr, _ := GetPostgresConnectionString("./secrets")

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		logger.Error("Unable to connect to database", "err", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: dbPool},
	}

	var greeting string
	err = dbPool.QueryRow(context.Background(), "SELECT 'Hello, Sir.'").Scan(&greeting)
	if err != nil {
		logger.Error("QueryRow failed", "err", err)
		os.Exit(1)
	}

	logger.Info("connection pool established")
	logger.Info("starting server", "addr", *addr)

	fmt.Println(greeting)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func GetPostgresConnectionString(secretsDir string) (string, error) {
	dbUser, err := os.ReadFile(filepath.Join(secretsDir, "app_user.txt"))
	if err != nil {
		return "", fmt.Errorf("failed to read db_user: %w", err)
	}

	dbUserPassword, err := os.ReadFile(filepath.Join(secretsDir, "app_password.txt"))
	if err != nil {
		return "", fmt.Errorf("failed to read db_password: %w", err)
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		strings.TrimSpace(string(dbUser)),
		strings.TrimSpace(string(dbUserPassword)),
		"snippetbox",
	)

	return connStr, nil
}
