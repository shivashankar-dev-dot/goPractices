package main

import (
	"fmt"

	aa "github.com/shivashankar-dev-dot/goPractices/values"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Print(add(4, 5))
	aa.CheckVals()
}
