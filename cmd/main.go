package main

import (
	"fmt"
	"os"

	"github.com/nunorocorreia/Go-Playground/cmd/hello"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}
