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
