package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"golang.org/x/exp/slog"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

func BadRequestError(w http.ResponseWriter, err error) {
	res := processError(err, w)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(res)
}

func InternalServerError(w http.ResponseWriter, err error) {
	res := processError(err, w)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(res)
}

func AuthorizationError(w http.ResponseWriter, err error) {
	res := processError(err, w)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(res)
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}


func processError(err error, w http.ResponseWriter) map[string]string {
	_, file, line, _ := runtime.Caller(1)
	logger.With("loc", fmt.Sprintf("%v:%v", file, line)).Error(err.Error())

	res := map[string]string{
		"error": err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	return res
}
