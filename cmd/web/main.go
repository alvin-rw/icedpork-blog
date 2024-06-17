package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alvin-rw/icedpork-blog/cmd/internal/data"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type config struct {
	port int
	db   struct {
		dsn                string
		maxOpenConns       int
		maxIdleConns       int
		maxIdleTimeSeconds int
	}
}

type application struct {
	config        config
	logger        *log.Logger
	templateCache map[string]*template.Template
	models        data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5003, "port")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "DB DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.IntVar(&cfg.db.maxIdleTimeSeconds, "db-max-idle-time-second", 900, "PostgreSQL max idle time in seconds")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	app := &application{
		config:        cfg,
		logger:        logger,
		templateCache: templateCache,
		models:        data.NewModels(db),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.router(),
	}

	logger.Printf("starting server on port %d", cfg.port)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	connMaxIdleTimeDuration := time.Duration(cfg.db.maxIdleTimeSeconds) * time.Second
	db.SetConnMaxIdleTime(connMaxIdleTimeDuration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
