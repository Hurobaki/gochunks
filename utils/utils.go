package utils

import (
	"path/filepath"
)

var Root = "."

func FullPath (path ...string) string {

	foo := filepath.Join(Root, filepath.Join(path...))

	return foo
}

