package utils

import (
	"encoding/json"
	"log"
	"net/http"

	cmn "github.com/prcryx/rss_aggregator/common"
	constants "github.com/prcryx/rss_aggregator/common/constants"
)

func ResponseWithJSONData(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf(cmn.FailedToMarshal)
		w.WriteHeader(constants.InternalServerErrCode)
		return
	}
	w.Header().Add(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(code)
	w.Write(data)
}
