package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// docker run --rm -d \
// --name dev-postgres \
// --network dev-network \
// -e POSTGRES_USER=postgres \
// -e POSTGRES_PASSWORD=password \
// -e POSTGRES_DB=postgres \
// -v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
// -p 5432:5432 \
// postgres

// NOTE: This is not ideal, because we will be establishing a new connection with a database
// every time we hit an endpoint. We would have to save connection pool somewhere and reuse
// for the whole duration of the application.
func PostgresConnection() (*pgxpool.Pool, error) {
	// Connect to postgres database using PGX
	const DATABASE_URL = "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse db config %v", err)
	}

	// TODO: Set config properties

	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create a databse pool %v", err)
	}

	dbConn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to acquire database connection %v", err)
	}

	if err := dbConn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database %v", err)
	}

	return dbPool, nil
}
