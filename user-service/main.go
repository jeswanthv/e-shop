package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeswanthv/e-shop/user-service/handlers"
	"github.com/jeswanthv/e-shop/user-service/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	handlers.InitDB(db)

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	s := r.PathPrefix("/api").Subrouter()
	s.Use(middleware.JWTAuth)
	s.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a protected route"))
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
