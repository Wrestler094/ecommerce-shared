package validation

import (
	"errors"

	v10 "github.com/go-playground/validator/v10"
)

type PlaygroundValidator struct {
	validate *v10.Validate
}

func NewPlaygroundValidator() *PlaygroundValidator {
	return &PlaygroundValidator{
		validate: v10.New(v10.WithRequiredStructEnabled()),
	}
}

func (v *PlaygroundValidator) Validate(s any) []FieldError {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	var ve v10.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]FieldError, 0, len(ve))
		for _, fe := range ve {
			out = append(out, FieldError{
				Field:   fe.Field(),
				Message: messageForTag(fe.Tag(), fe.Param()),
			})
		}
		return out
	}

	return nil
}
