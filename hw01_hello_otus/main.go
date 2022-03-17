package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	result := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(result)
}
