package utils

import (
	"encoding/json"
	"log"
	"net/http"

	constants "github.com/prcryx/rss_aggregator/common/constants"
	errs "github.com/prcryx/rss_aggregator/common/errs"
)

func ResponseWithJSONData(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf(errs.FailedToMarshal)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(code)
	w.Write(data)
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("5xx Error: %v", msg)
	}

	ResponseWithJSONData(w, code, errs.ErrorResponse{
		Error: msg,
	})
}
