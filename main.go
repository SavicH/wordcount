package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	res := len(strings.Split(input, " "))
	fmt.Print(res)
}
