package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"snippetbox.gentiluomo.dev/internal/models"
)

type application struct {
	logger   *slog.Logger
	cache    *redis.Client
	db       *pgxpool.Pool
	snippets *models.SnippetModel
	users    *models.UserModel
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

	var greeting string
	err = dbPool.QueryRow(context.Background(), "SELECT 'Hello, Sir.'").Scan(&greeting)
	if err != nil {
		logger.Error("QueryRow failed", "err", err)
		os.Exit(1)
	}
	logger.Info("connection pool established")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err = redisClient.Ping(context.Background()).Err()
	if err != nil {
		logger.Error("Unable to connect to redis", "err", err)
		os.Exit(1)
	}
	logger.Info("connection to redis established")

	app := &application{
		logger:   logger,
		cache:    redisClient,
		db:       dbPool,
		snippets: &models.SnippetModel{DB: dbPool},
		users:    &models.UserModel{DB: dbPool},
	}

	tlsConfig := &tls.Config{
		// use only elliptic curves with assembly implementations
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", "addr", *addr)

	err = srv.ListenAndServeTLS("./tls/certs-chain.pem", "./tls/key.pem")
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
