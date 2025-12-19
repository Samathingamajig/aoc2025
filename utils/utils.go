package utils

import (
	"fmt"
	"os"
	"path"
)

func ReadFilePanic(name ...string) string {
	raw, err := os.ReadFile(path.Join(name...))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(raw)
}

func Bool2int(b bool) int {
	// https://dev.to/chigbeef_77/bool-int-but-stupid-in-go-3jb3
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}
