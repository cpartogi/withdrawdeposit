package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cpartogi/withdrawdeposit/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func calculateMin(fieldName, num string) string {
	fieldsName := map[string]int{
		"phone number": 2,
	}

	i, err := strconv.Atoi(num)
	if err != nil {
		return num
	}

	for key := range fieldsName {
		if key == fieldName {
			return strconv.Itoa(i - fieldsName[key])
		}
	}

	return num
}

func oneofErrorSplit(choices string) (result string) {
	s := strings.Split(choices, " ")
	for idx := range s {
		message := fmt.Sprintf("%s or ", s[idx])
		if (idx + 1) == len(s) {
			message = s[idx]
		}
		result += message
	}
	return
}

func errorType(err error) (int, error) {
	switch {
	case isPqError(err):
		return helper.PqError(err)
	}
	return helper.CommonError(err)
}

func isPqError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}
	return false
}

func switchErrorValidation(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			field := SetLowerAndAddSpace(err.Field())

			// Change Field Name
			switch field {
			case "msisdn":
				field = "phone number"
			case "otp":
				field = "otp code"
			}

			// Check Error Type
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is mandatory",
					field)
			case "email":
				message = fmt.Sprintf("%s is not valid email",
					field)
			case "number":
				message = fmt.Sprintf("%s must be numbers only",
					field)
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					field, err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					field, err.Param())
			case "min":
				minimum := calculateMin(field, err.Param())
				message = fmt.Sprintf("%s at least %s characters long",
					field, minimum)
			case "max":
				message = fmt.Sprintf("the length of %s must be %s characters or fewer",
					field, err.Param())
			case "startswith":
				message = fmt.Sprintf("%s must starts with %s",
					field, err.Param())
			case "len":
				message = fmt.Sprintf("%s length must %s characters",
					field, err.Param())
			case "oneof":
				choices := oneofErrorSplit(err.Param())
				message = fmt.Sprintf("%s must specify one of %s",
					field, choices)
			default:
				message = err.Error()
			}

			break
		}
	}
	return
}
