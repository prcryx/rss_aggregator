package main

import (
	"fmt"
	"log"
	"net/http"
	config "github.com/prcryx/rss_aggregator/config"
	routes "github.com/prcryx/rss_aggregator/internal/routes"
)

func main() {

	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	run(env)

}

func run(env config.EnvVars) {
	rootRoute := routes.Root()
	routes.MountAll(rootRoute, routes.V1)

	//create server
	srv := http.Server{
		Handler: rootRoute,
		Addr:    ":" + env.Port,
	}

	fmt.Printf("server is starting on port: %v", env.Port)

	//listen
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
