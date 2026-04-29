// Package goose provides database migration tooling.
// It is a fork of pressly/goose with additional features and improvements.
package goose

import (
	"context"
	"database/sql"
	"fmt"
)

// Version is the current version of goose.
const Version = "v3.0.0"

// Dialect represents a supported database dialect.
type Dialect string

const (
	DialectPostgres   Dialect = "postgres"
	DialectMySQL      Dialect = "mysql"
	DialectSQLite3    Dialect = "sqlite3"
	DialectMSSQL      Dialect = "mssql"
	DialectRedshift   Dialect = "redshift"
	DialectTiDB       Dialect = "tidb"
	DialectClickHouse Dialect = "clickhouse"
	DialectVertica    Dialect = "vertica"
)

// ErrNoCurrentVersion is returned when there is no current migration version.
var ErrNoCurrentVersion = fmt.Errorf("no current version found")

// ErrNoNextVersion is returned when there is no next migration version.
var ErrNoNextVersion = fmt.Errorf("no next version found")

// Provider manages database migrations.
type Provider struct {
	db      *sql.DB
	dialect Dialect
	dir     string
	opts    Options
}

// Options holds configuration options for the migration provider.
type Options struct {
	// TableName is the name of the migrations table. Defaults to "goose_db_version".
	TableName string
	// NoVersioning disables versioning and runs migrations without tracking.
	NoVersioning bool
	// AllowMissing allows missing (out-of-order) migrations to be applied.
	AllowMissing bool
	// Verbose enables verbose logging.
	Verbose bool
}

// MigrationResult holds the result of a single migration operation.
type MigrationResult struct {
	Source    MigrationSource
	Direction string
	Duration  int64 // duration in milliseconds
	Empty     bool  // true if the migration was a no-op (e.g., no SQL statements)
}

// MigrationSource describes the source of a migration file.
type MigrationSource struct {
	Type    string // "sql" or "go"
	Path    string
	Version int64
}

// NewProvider creates a new migration provider.
//
// The db parameter is the database connection to use.
// The dialect parameter specifies the database dialect.
// The dir parameter is the directory containing migration files.
func NewProvider(dialect Dialect, db *sql.DB, dir string, opts Options) (*Provider, error) {
	if db == nil {
		return nil, fmt.Errorf("db must not be nil")
	}
	if dir == "" {
		return nil, fmt.Errorf("migration directory must not be empty")
	}
	if opts.TableName == "" {
		opts.TableName = "goose_db_version"
	}
	return &Provider{
		db:      db,
		dialect: dialect,
		dir:     dir,
		opts:    opts,
	}, nil
}

// Up applies all pending migrations.
func (p *Provider) Up(ctx context.Context) ([]*MigrationResult, error) {
	return p.up(ctx, false, 0)
}

// UpByOne applies the next pending migration.
func (p *Provider) UpByOne(ctx context.Context) (*MigrationResult, error) {
	results, err := p.up(ctx, true, 0)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, ErrNoNextVersion
	}
	return results[0], nil
}

// UpTo applies pending migrations up to and including the specified version.
func (p *Provider) UpTo(ctx context.Context, version int64) ([]*MigrationResult, error) {
	return p.up(ctx, false, version)
}

// Down rolls back the most recently applied migration.
func (p *Provider) Down(ctx context.Context) (*MigrationResult, error) {
	return p.down(ctx, false, 0)
}

// DownTo rolls back all migrations down to and including the specified version.
func (p *Provider) DownTo(ctx context.Context, version int64) (*MigrationResult, error) {
	return p.down(ctx, false, version)
}

// Status returns the status of all migrations.
func (p *Provider) Status(ctx context.Context) ([]*MigrationStatus, error) {
	return nil, fmt.Errorf("not implemented")
}

// MigrationStatus represents the status of a single migration.
type MigrationStatus struct {
	Source    MigrationSource
	AppliedAt string // ISO 8601 timestamp or empty if not applied
}

func (p *Provider) up(ctx context.Context, byOne bool, version int64) ([]*MigrationResult, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
}

func (p *Provider) down(ctx context.Context, byOne bool, version int64) (*MigrationResult, error) {
	// TODO: implement
	return nil, fmt.Errorf("not implemented")
}
