package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	//"log"
	"net/http"
	"strconv"
)
var (
	ErrorBadRequest = errors.New("invalid request parameter")
)
// decodeArithmeticRequest decode request params to struct
func decodeArithmeticRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("decodeArithmeticRequest");
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequest
	}

	a, _ := strconv.Atoi(pa)
	b, _ := strconv.Atoi(pb)
fmt.Println(a);
	fmt.Println(b);
	return ArithmeticRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil
}

// encodeArithmeticResponse encode response to return
func encodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}




	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(httptransport.NewServer(
		endpoint,
		decodeArithmeticRequest,
		encodeArithmeticResponse,
		options...,
	))

	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
func codeFrom(err error) int {
	switch err {
	//case ErrNotFound:
	//	return http.StatusNotFound
	//case ErrAlreadyExists, ErrInconsistentIDs:
	//	return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}