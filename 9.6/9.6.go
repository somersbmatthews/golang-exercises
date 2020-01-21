package main

import (
	"fmt"
	"runtime"
)

func main() {
	FindArch()
}

func FindArch() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
