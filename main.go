package main

import (
	"fmt"
	"runtime"
)

func goversion() string {
	return runtime.Version()
}

func main() {
	fmt.Println("Your Go runtime version:", goversion())
}
