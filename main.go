package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Muhamaddiis/Task-Management-Api/database"
	"github.com/Muhamaddiis/Task-Management-Api/handlers"
	"github.com/Muhamaddiis/Task-Management-Api/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()
	database.Migrate()

	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.ContentTypeMiddleware)

	api.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	api.HandleFunc("/tasks", handlers.GetAllTasks).Methods("GET")

	//start server
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("Server starting on port :%v\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
