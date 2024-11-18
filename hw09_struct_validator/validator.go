package hw09structvalidator

func Validate(v interface{}) ValidationErrors {
	return NewValidator().Validate(v)
}
