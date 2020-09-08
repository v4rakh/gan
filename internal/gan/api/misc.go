package api

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/v4rakh/gan/internal/gan/api/presenter"
	"github.com/v4rakh/gan/internal/util"
)

func ConvertErrorsTo(val *validator.ValidationErrors) *presenter.ErrorResponse {
	errorMap := make(map[string]string)
	for _, v := range *val {
		key, txt := validatorErrorToText(&v)
		errorMap[key] = txt
	}

	return presenter.NewErrorResponseWithStatusAndMessageAndData(presenter.ErrorBadRequest, util.ValuesString(errorMap), errorMap)
}

func validatorErrorToText(e *validator.FieldError) (string, string) {
	x := *e

	switch x.Tag() {
	case "required":
		return x.Field(), fmt.Sprintf("%s is required", x.Field())
	case "max":
		return x.Field(), fmt.Sprintf("%s cannot be longer than %s", x.Field(), x.Param())
	case "min":
		return x.Field(), fmt.Sprintf("%s must be longer than %s", x.Field(), x.Param())
	case "len":
		return x.Field(), fmt.Sprintf("%s must be %s characters long", x.Field(), x.Param())
	}
	return x.Field(), fmt.Sprintf("%s is not valid", x.Field())
}
