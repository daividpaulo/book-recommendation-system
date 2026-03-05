package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	httpdelivery "book-recommendation-system/recommendations-api/internal/delivery/http"
	mlservice "book-recommendation-system/recommendations-api/internal/repository/mlservice"
	postgresrepo "book-recommendation-system/recommendations-api/internal/repository/postgres"
	"book-recommendation-system/recommendations-api/internal/usecase"
)

func main() {
	port := getEnv("PORT", "8080")
	db := mustConnectDB()
	mlURL := getEnv("ML_SERVICE_URL", "http://ml-recommendations-api:5000")

	postgresRepository := postgresrepo.New(db)
	mlGateway := mlservice.New(mlURL)
	service := usecase.New(postgresRepository, postgresRepository, postgresRepository, mlGateway)
	handlers := httpdelivery.NewHandlers(service)
	router := httpdelivery.NewRouter(handlers)

	addr := ":" + port
	log.Printf("recommendations-api listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func mustConnectDB() *sql.DB {
	host := getEnv("DB_HOST", "postgres")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "book_recommendations")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

	for attempt := 1; attempt <= 20; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		err = database.PingContext(ctx)
		cancel()
		if err == nil {
			log.Println("connected to postgres")
			return database
		}
		log.Printf("postgres not ready (attempt %d/20): %v", attempt, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatalf("could not connect to postgres: %v", err)
	return nil
}
