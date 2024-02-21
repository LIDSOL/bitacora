// cli/input.go
package cli

import (
	"regexp"
	"strings"
)

func CleanString(s string) string {
	r := strings.TrimSpace(s)
	r = strings.ToUpper(r)
	return r
}

func IsValidRFC(s string) bool {
	if len(s) != 13 {
		return false
	}

	reg := `^[A-Z]{4}[0-9]{2}(0[1-9]|1[0-2])(0[1-9]|1[0-9]|2[0-9]|3[0-1])[A-Z0-9]{3}$`

	m, _ := regexp.MatchString(reg, s)
	return m
}

func IsValidAccountNumber(s string) bool {
	if len(s) != 9 {
		return false
	}
	reg := `^[0-9]{9}$`

	m, _ := regexp.MatchString(reg, s)
	return m
}

func IsValidEmail(s string) bool {
	if len(s) > 256 {
		return false
	}

	reg := `.+@.+\..+`

	m, _ := regexp.MatchString(reg, s)
	return m
}

func IsValidName(s string) bool {
	if len(s) > 128 {
		return false
	}

	reg := `^[a-zA-Z áéíóúÁÉÍÓÚ]+$`

	m, _ := regexp.MatchString(reg, s)
	return m
}
