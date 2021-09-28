package http

import (
	"net/http"
)

func homePageHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	_, _ = rw.Write([]byte("siema"))

}
