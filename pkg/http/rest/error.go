package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Err struct {
	ErrorText string `json:"errorText"`
	Message   string `json:"message"`
}

func LogError(r *http.Request, status int, err error) {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	logger.Error(err.Error(),
		"Method", r.Method,
		"Status", status,
		"URL", r.URL.String(),
	)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, err error) {
	data := Err{
		ErrorText: http.StatusText(status),
		Message:   err.Error(),
	}

	w.WriteHeader(status)
	LogError(r, status, err)
	Encode(w, &data)
}

func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	LogError(r, http.StatusInternalServerError, err)
	err = fmt.Errorf("the server encountered a problem and could not process your request")
	ErrorResponse(w, r, http.StatusInternalServerError, err)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("the requested resource could not be found")
	ErrorResponse(w, r, http.StatusNotFound, err)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("the %s method is not supported for this resource", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, err)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err)
}
