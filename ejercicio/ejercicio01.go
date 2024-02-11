package ejercicio

import "strconv"

func Ejercicio01(texto string) (int, string) {
	num, err := strconv.Atoi(texto)

	if err != nil {
		return 0, "Error al convertir el texto"
	}
	if num > 100 {
		return num, "El numero es mayor a 100"
	} else {
		return num, "El numero es menor a 100"
	}
}
