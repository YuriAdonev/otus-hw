package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	result := reverse.String("Hello, OTUS!")
	fmt.Println(result)
}
