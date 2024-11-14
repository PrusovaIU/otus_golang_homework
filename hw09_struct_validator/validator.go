package main

func Validate(v interface{}) ValidationErrors {
	return NewValidator().Validate(v)
}
