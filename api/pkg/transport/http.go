package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func NewHttpService(endpoints Endpoints, logger log.Logger) *mux.Router {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r.Methods(http.MethodPost).Path("/api/v1/news").Handler(kithttp.NewServer(
		endpoints.CreateNews,
		decodeCreateNewsRequest,
		encodeResponse,
		options...,
	))
	r.Methods(http.MethodGet).Path("/api/v1/news/{id}").Handler(kithttp.NewServer(
		endpoints.GetNewsById,
		decodeGetNewById,
		encodeResponse,
		options...,
	))
	return r
}

func decodeCreateNewsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateNewsRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetNewById(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id := vars["id"]
	newsId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return GetNewsByIdRequest{
		Id: uint64(newsId),
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(e error) int {
	switch e {
	// record not found return 404
	default:
		return http.StatusInternalServerError
	}
}
