package validation

// FieldError — структура для описания ошибок валидации.
type FieldError struct {
	Field   string
	Message string
}

// Validator — интерфейс для адаптируемых валидаторов.
type Validator interface {
	Validate(i any) []FieldError
}
