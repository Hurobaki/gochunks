package format

import "strings"

func Concatenate (separator string, args ...string) string {
	return strings.Join(args, separator)
}
