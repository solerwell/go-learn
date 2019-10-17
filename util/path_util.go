package util

import (
	"fmt"
	"os"
)

func GetProjectRootPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return path
}
