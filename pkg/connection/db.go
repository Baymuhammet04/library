package connection

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2004942/library/internal/config"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Ensure Database implements DB and DBOps.
var _ DB = (*Database)(nil)

// Querier defines basic query operations.
type Querier interface {
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

// DB extends Querier with additional convenience methods.
type DB interface {
	Querier
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type Database struct {
	Pool *pgxpool.Pool
}

func NewDBConnection(ctx context.Context, cfg config.PostgresConfig) (*Database, error) {
	const (
		retryAttempts = 3
		retryDelay    = 2 * time.Second
	)

	var (
		db  *pgxpool.Pool
		err error
	)

	for i := 0; i < retryAttempts; i++ {
		db, err = connectToDB(ctx, cfg)
		if err == nil {
			return &Database{Pool: db}, nil
		}

		time.Sleep(retryDelay)
	}

	// After three times, if db connection failed, then throw fatal.
	log.Fatalf("failed to connect to db after %d attempts: %v", retryAttempts, err)

	return nil, fmt.Errorf("failed to connect to db after %d attempts: %w", retryAttempts, err)
}

func connectToDB(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.GenerateDSN())
	if err != nil {
		return nil, fmt.Errorf("parsing connection config: %w", err)
	}

	// create a new connection pool.
	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	// ping the database to verify the connection.
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pinging database error: %w", err)
	}

	return db, nil
}

// QueryRow func is used for querying single row.
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.Pool.QueryRow(ctx, query, args...)
}

// Query func is used for querying multiple rows.
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.Pool.Query(ctx, query, args...)
}

// Exec func is used for executing query.
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	result, err := d.Pool.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("executing query error: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgconn.CommandTag{}, pgx.ErrNoRows
	}

	return result, nil
}

// Get func is used for getting single row.
func (d *Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, d.Pool, dest, query, args...)
}

// Select func is used for getting multiple rows.
func (d *Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, d.Pool, dest, query, args...)
}

