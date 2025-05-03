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
	"github.com/redis/go-redis/v9"
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
	fmt.Println(greeting)

	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err = rdb.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(context.Background(), "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	logger.Info("connection to redis established")

	logger.Info("starting server", "addr", *addr)

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
