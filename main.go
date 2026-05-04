package main

import (
	"fmt"
	"log"
	"net/http"

	"f1-tracker-backend/internal/config"
	"f1-tracker-backend/internal/database"
	"f1-tracker-backend/internal/router"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DSN())
	defer db.Close()

	r := router.New(db)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🏎️  F1 Tracker API running on http://localhost%s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
