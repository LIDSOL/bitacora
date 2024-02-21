// cli/input.go
package cli

import (
	"strings"
)

func CleanString(s string) string {
	r := strings.TrimSpace(s)
	r = strings.ToUpper(r)
	return r
}
