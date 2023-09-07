package handlers

import (
	"net/http"

	errs "github.com/prcryx/rss_aggregator/common/errs"
	utils "github.com/prcryx/rss_aggregator/common/utils"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request) {
	utils.ResponseWithError(w, http.StatusBadRequest, errs.SomethingWentWrong)
}
