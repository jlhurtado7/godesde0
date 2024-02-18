package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {
	//variables.MuestroEnteros()

	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(2345)
	fmt.Println(estado)
	fmt.Println(texto)
	/*
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
			fmt.Println("Este si es", os)
		}

		numero, texto := ejercicio.Ejercicio01("iiogol")
		fmt.Println(numero)
		fmt.Println(texto)

		sum := sum(os.Args[1], os.Args[2])
		fmt.Println("Sum:", sum)
	*/
}

/*
func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}
*/
