package helper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseError(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			for _, e := range typedError {
				out = append(out, ParseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			out = append(out, ParseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}

	return out
}

func ParseFieldError(e validator.FieldError) string {
	fieldPrefix := fmt.Sprintf("The field %s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", fieldPrefix)
	case "email":
		return fmt.Sprintf("%s must be an email", fieldPrefix)
	default:
		return fmt.Errorf("%v", e).Error()
	}
}

func ParseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("The field %s must be a %s", e.Field, e.Type.String())
}
