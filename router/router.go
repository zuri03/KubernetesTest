package router

import (
	"encoding/json"
	"net/http"
	"sync"
)

type APIRouter struct {
	waitgroup *sync.WaitGroup
}

type ServerResponse struct {
	Data string `json:"data"`
}

func New(waitgroup *sync.WaitGroup) *http.ServeMux {
	router := http.NewServeMux()
	apiRouter := APIRouter{waitgroup: waitgroup}
	router.HandleFunc("/api", apiRouter.serveHTTP)
	return router
}

func (router *APIRouter) serveHTTP(writer http.ResponseWriter, request *http.Request) {
	router.waitgroup.Add(1)
	defer router.waitgroup.Done()

	successResponse := &ServerResponse{}

	switch request.Method {
	case http.MethodGet:
		successResponse.Data = "GET: success"
		jsonBytes, err := json.Marshal(successResponse)
		if err != nil {
			http.Error(writer, "Internal server error, unable to marshal json", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(200)
		writer.Write(jsonBytes)
		return
	case http.MethodPost:
		successResponse.Data = "POST: success"
		jsonBytes, err := json.Marshal(successResponse)
		if err != nil {
			http.Error(writer, "Internal server error, unable to marshal json", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(200)
		writer.Write(jsonBytes)
		return
	case http.MethodDelete:
		successResponse.Data = "DELETE: success"
		jsonBytes, err := json.Marshal(successResponse)
		if err != nil {
			http.Error(writer, "Internal server error, unable to marshal json", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(200)
		writer.Write(jsonBytes)
		return
	case http.MethodPut:
		successResponse.Data = "PUT: success"
		jsonBytes, err := json.Marshal(successResponse)
		if err != nil {
			http.Error(writer, "Internal server error, unable to marshal json", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(200)
		writer.Write(jsonBytes)
		return
	default:
		http.Error(writer, "Method not supported", http.StatusMethodNotAllowed)
	}
}
