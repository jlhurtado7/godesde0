package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var stado bool
var Sueldo float32
var Fechas time.Time

func RestoVariables() {

	Nombre = "Joel"
	stado = true
	Sueldo = 1000.50
	Fechas = time.Now()

	fmt.Println("Nombre:", Nombre)
	fmt.Println("Estado:", stado)
	fmt.Println("Sueldo:", Sueldo)
	fmt.Println("Fechas:", Fechas)
}

func ConviertoTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto

}
