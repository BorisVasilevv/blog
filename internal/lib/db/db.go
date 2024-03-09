package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const connection = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s application_name=%s"

type Db struct {
	PgConn *pgxpool.Pool
}

func New(ctx context.Context) *Db {

	connectionString := fmt.Sprintf(connection, "localhost", "5432", "postgres", "admin", "postgres", "disable")

	conn, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return &Db{PgConn: conn}
}
