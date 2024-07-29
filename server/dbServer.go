package server

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func InitDatabase(config *viper.Viper) *pgxpool.Conn {
	connectionString := config.GetString("database.connection_string")
	maxIdleConnections := config.GetInt("database.max_idle_connections")
	maxOpenConnections := config.GetInt("database.max_open_connections")
	connectionMaxLifetime := config.GetDuration("database.connection_max_lifetime")
	//driverName := config.GetString("database.driver_name")

	if connectionString == "" {
		log.Fatalf("Database connection string is missing")
	}

	dbConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = int32(maxOpenConnections)
	dbConfig.MaxConnIdleTime = time.Duration(maxIdleConnections)
	dbConfig.MaxConnLifetime = connectionMaxLifetime

	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer log.Println("Database already connected")
	return connection
}
