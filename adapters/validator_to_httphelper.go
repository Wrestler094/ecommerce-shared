package adapters

import (
	"github.com/Wrestler094/ecommerce-shared/httphelper"
	"github.com/Wrestler094/ecommerce-shared/validation"
)

// HttpValidatorAdapter адаптирует интерфейс validation.Validator к интерфейсу httphelper.Validator.
type HttpValidatorAdapter struct {
	inner validation.Validator
}

// NewHttpValidatorAdapter возвращает адаптированную реализацию httphelper.Validator
func NewHttpValidatorAdapter(inner validation.Validator) httphelper.Validator {
	return &HttpValidatorAdapter{inner}
}

// Validate адаптирует ошибки из validation.FieldError в httphelper.FieldError.
func (a *HttpValidatorAdapter) Validate(i any) []httphelper.FieldError {
	src := a.inner.Validate(i)
	out := make([]httphelper.FieldError, len(src))

	for i, err := range src {
		out[i] = httphelper.FieldError{
			Field:   err.Field(),
			Message: err.Message(),
		}
	}

	return out
}
