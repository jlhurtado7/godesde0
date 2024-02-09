package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoTexto(2345)
	fmt.Println(estado)
	fmt.Println(texto)
}
