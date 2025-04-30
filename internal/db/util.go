package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

// TODO: consider using pgxpool.Connect(...)
func ConnectDB() *pgx.Conn {
	dbPort := 5432
	host := "localhost"
	username := "root"
	dbName := "postgres"
	pwd := "lab123"

	connectionCmd := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, pwd, host, dbPort, dbName)
	log.Printf(connectionCmd)
	conn, err := pgx.Connect(context.Background(), connectionCmd)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Printf("Connected to database successfully")
	return conn
}

func CloseDB(conn *pgx.Conn) {
	err := conn.Close(context.Background())
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
	log.Printf("Database connection closed successfully")

}

// ticker, company name, sector, industry, market cap
func CreateStocksTable(dbConnection *pgx.Conn) {
	command := `
	CREATE TABLE IF NOT EXISTS stocks (
		ticker VARCHAR(5) PRIMARY KEY,
		company_name VARCHAR(30),
		sector VARCHAR(30),
		industry VARCHAR(30),
		market_cap DECIMAL(10, 2)
		);
	`
	_, err := dbConnection.Exec(context.Background(), command)
	if err != nil {
		log.Fatal("Failed to create stocks table: ", err)
	}
	log.Printf("Created stocks table successfully")
}
