package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Load(cfg *Config) (*sql.DB, error) {
	const op = "postgres.Load"

	dst := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.User,
		cfg.Password,
	)

	db, err := sql.Open("postgres", dst)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("Connected to database!\n")
	return db, nil
}

func MustLoad(cfg *Config) *sql.DB {
	db, err := Load(cfg)
	if err != nil {
		panic(err)
	}
	return db
}

func MigrateUp(db *sql.DB) error {
	const op = "postgres.MigrateUp"

	driver, err := pg.WithInstance(db, &pg.Config{})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://migration", "postgres", driver)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func MustMigrateUp(db *sql.DB) {
	if err := MigrateUp(db); err != nil {
		panic(err)
	}
}
