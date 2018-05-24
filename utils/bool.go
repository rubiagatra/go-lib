package utils

import (
	"regexp"
	"strconv"
)

// IsEmailValid -> validate email using regex
func IsEmailValid(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

// IsNumeric -> Check if input string is int
func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)

	if err != nil {
		return false
	}

	return true
}

// Bool2String returns true or false depending on input
func Bool2String(b bool) string {
	return strconv.FormatBool(b)
}

// IsStringInSlice check if string exists in slice
func IsStringInSlice(s string, list []string) bool {
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}
