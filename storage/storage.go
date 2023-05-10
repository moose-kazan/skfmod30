package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBPool struct {
	db  *pgxpool.Pool
	ctx context.Context
}

type DBPoolIface interface {
	Ping()
}

func (d *DBPool) Ping() {
	d.db.Ping(d.ctx)
}

func New(dsn string) *DBPool {
	var d DBPool
	var err error
	d.ctx = context.Background()
	d.db, err = pgxpool.Connect(d.ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &d
}
