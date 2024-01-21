package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KunalRajpal/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
    r := mux.NewRouter()

    //Registering the bookstore routes
    routes.RegisterBookStoreRoutes(r)
    http.Handle("/", r)

    // Serve static files from the 'web' directory
    r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("web/"))))

    // Use the PORT environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "9010" // Default port if not specified
    }

    log.Fatal(http.ListenAndServe(":" + port, r))
}