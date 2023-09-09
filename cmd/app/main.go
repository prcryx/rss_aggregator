package main

import (
	"fmt"
	"log"
	"net/http"

	chi "github.com/go-chi/chi"
	"github.com/prcryx/rss_aggregator/cmd/routes"
	config "github.com/prcryx/rss_aggregator/config"
)

func main() {

	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	run(env)

}

func run(env *config.EnvVars) {
	rootRoute := routes.Root()
	routes.MountAll(rootRoute, routes.V1)

	//create server
	srv := CreateServer(env, rootRoute)

	fmt.Printf("server is starting on port: %v", env.Port)

	//listen
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func CreateServer(env *config.EnvVars, root chi.Router) *http.Server {
	return &http.Server{
		Handler: root,
		Addr:    ":" + env.Port,
	}
}
