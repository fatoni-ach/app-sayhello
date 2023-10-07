package main

import (
	"fmt"

	"github.com/fatoni-ach/app-sayhello/helpers"
	sayhello "github.com/fatoni-ach/go-module-sayhello/v2"
)

func main() {
	fmt.Println(sayhello.Hello("Anonim"))

	fmt.Println("Jumlah 2 + 2 = ", helpers.Sum(2, 2))
}
