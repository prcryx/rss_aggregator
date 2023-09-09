package routes

import (
	chi "github.com/go-chi/chi"
	cors "github.com/go-chi/cors"
	handlers "github.com/prcryx/rss_aggregator/handlers"
)

func Root() chi.Router {
	rootRoute := chi.NewRouter()

	//middleware
	rootRoute.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowCredentials: true,
				MaxAge:           300,
				ExposedHeaders:   []string{"Link"},
			},
		),
	)

	return rootRoute
}

func MountAll(root chi.Router, mountPath string) {
	router := chi.NewRouter()

	// all handlers here
	router.Get(HealthCheck, handlers.HealthCheckHandler)
	router.Get(Error, handlers.ErrorHandler)

	root.Mount(mountPath, router)
}
