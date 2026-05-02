package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error abriendo base de datos: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	log.Println("Database conectada exitosamente")
	return db
}
