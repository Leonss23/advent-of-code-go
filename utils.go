package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetInput() string {
	filename := "input"
	_, callerPath, _, _ := runtime.Caller(1)

	path, err := filepath.Abs(callerPath + "/../" + filename)
	check(err)
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic("xd")
	}
}
