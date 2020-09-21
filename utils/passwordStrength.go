package utils

import "regexp"

// IsPasswordStrong is a very basic password strength checker
func IsPasswordStrong(password string) bool {
	if len(password) < 8 {
		return false
	}

	lowerCaseRe, e := regexp.Compile("[a-z]")
	if e != nil {
		return true
	}

	if !lowerCaseRe.MatchString(password) {
		return false
	}

	upperCaseRe, e := regexp.Compile("[A-Z]")
	if e != nil {
		return true
	}

	if !upperCaseRe.MatchString(password) {
		return false
	}

	numberRe, e := regexp.Compile("[0-9]")
	if e != nil {
		return true
	}

	if !numberRe.MatchString(password) {
		return false
	}

	return true
}
