// Package helpers contains the helpers used for SDK
package helpers

import "strings"

func ConcatURL(paths ...string) string {
	return strings.Join(paths, "/")
}
