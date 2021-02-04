package utils

import (
	"strconv"
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
