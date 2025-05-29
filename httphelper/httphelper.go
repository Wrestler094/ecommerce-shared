package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/Wrestler094/ecommerce-shared/validation"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details any    `json:"details,omitempty"`
}

type FieldErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func RespondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, ErrorResponse{
		Error: message,
	})
}

func RespondValidationErrors(w http.ResponseWriter, errs []validation.FieldError) {
	if len(errs) == 0 {
		RespondError(w, http.StatusUnprocessableEntity, "Invalid data")
		return
	}

	out := make([]FieldErrorResponse, 0, len(errs))
	for _, e := range errs {
		out = append(out, FieldErrorResponse{
			Field:   e.Field,
			Message: e.Message,
		})
	}

	RespondJSON(w, http.StatusUnprocessableEntity, ErrorResponse{
		Error:   "Invalid data",
		Details: out,
	})
}

func DecodeAndRespond[T any](r *http.Request, w http.ResponseWriter, validator validation.Validator) (T, bool) {
	var req T
	var zero T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid JSON")
		return zero, false
	}

	if errs := validator.Validate(req); errs != nil {
		RespondValidationErrors(w, errs)
		return zero, false
	}

	return req, true
}
