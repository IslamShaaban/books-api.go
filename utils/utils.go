package utils

import "regexp"

func RemoveSpecialChars(str string) string {
	// Remove non-alphanumeric characters using a regular expression
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(str, "")
}
