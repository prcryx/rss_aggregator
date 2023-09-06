package routes

import (
	"github.com/go-chi/chi"
	"github.com/prcryx/rss_aggregator/handlers"
)



func MountAll(root chi.Router, mountPath string){
    router := chi.NewRouter()

    // all handlers here
    router.Get(HealthCheck, handlers.HealthCheck)

    root.Mount(mountPath, router)
}