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

	path, err := filepath.Abs(callerPath + "/../../" + filename)
	Check(err)
	data, err := os.ReadFile(path)
	Check(err)
	return string(data)
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		panic("xd")
	}
}

func Sum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}
