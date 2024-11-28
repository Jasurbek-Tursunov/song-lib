package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func Load(cfg *Config) (*sql.DB, error) {
	const op = "postgres."
	dst := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.User,
		cfg.Password,
	)

	db, err := sql.Open("postgres", dst)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Connected to database ...")
	return db, nil
}

func MustLoad(cfg *Config) *sql.DB {
	db, err := Load(cfg)
	if err != nil {
		panic(err)
	}
	return db
}
