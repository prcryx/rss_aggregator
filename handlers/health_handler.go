package handlers

import (
	"net/http"

	utils "github.com/prcryx/rss_aggregator/common/utils"
)

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	utils.ResponseWithJSONData(w, http.StatusOK, struct{}{})
}
