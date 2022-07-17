package app

import (
	"encoding/json"
	"github.com/Vladimir77715/otus/k8s_base/internal/models"
	"net/http"
)

type response struct {
	Status string `json:"status"`
}

var (
	okResponse, _ = json.Marshal(response{Status: "OK"})
	headers       = []models.Header{models.ContentTypeJson, models.CharsetUft8}
)

func Heath(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		for _, v := range headers {
			writer.Header().Add(v[0], v[1])
		}
		if _, err := writer.Write(okResponse); err != nil {
			println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}
