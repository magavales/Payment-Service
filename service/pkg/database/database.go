package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"service/pkg/database/tables"
)

type Database struct {
	Pool   *pgxpool.Pool
	Access tables.DataAccess
}

func (db *Database) Connect() {
	poolConn, _ := pgxpool.ParseConfig("user=postgres password=1703 host=localhost port=5432 dbname=postgres pool_max_conns=10")

	var err error
	db.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConn)
	if err != nil {
		log.Printf("I can't connect to database: %s\n", err)
	}
}
