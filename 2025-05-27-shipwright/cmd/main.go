package main

import (
	"fmt"
)

var buildSha string

func main() {
	fmt.Printf("Hello cloud native bergen community!\n")
	fmt.Printf("Build=%s\n", buildSha)
}
