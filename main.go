package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	chi "github.com/go-chi/chi"
	"github.com/go-chi/cors"
	godotenv "github.com/joho/godotenv"
	cmn "github.com/prcryx/rss_aggregator/common"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(cmn.PortNotFound)
	}

	router := chi.NewRouter()

	//middleware
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedHeaders:   []string{"*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowCredentials: false,
				MaxAge:           300,
				ExposedHeaders:   []string{"Link"},
			},
		),
	)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("listening on port: %v...", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}