package main

import (
	"fmt"

	"github.com/godesde0/ejercicio"
)

func main() {
	//variables.MuestroEnteros()
	/*
		variables.RestoVariables()
		estado, texto := variables.ConviertoaTexto(2345)
		fmt.Println(estado)
		fmt.Println(texto)

		if os := runtime.GOOS; os == "linux" || os == "Mac" {
			fmt.Println("Este sistema no es", os)
		} else {
			fmt.Println("Este si es", os)
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Este sistema no es", os, "esto es linux")
		case "Mac":
			fmt.Println("Este sistema no es", os, "esto es Mac")
		default:
			fmt.Printf("%s \n", os)
		}
	*/
	numero, texto := ejercicio.Ejercicio01("90")
	fmt.Println(numero)
	fmt.Println(texto)

}
