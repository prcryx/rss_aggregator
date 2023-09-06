package handlers

import (
	"net/http"

	utils "github.com/prcryx/rss_aggregator/common/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.ResponseWithJSONData(w, 200, struct{}{})
}
