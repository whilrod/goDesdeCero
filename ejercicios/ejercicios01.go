package ejercicios

import "strconv"

func DevuelveDosValores(numero string) (int, string) {
	texto := ""
	numeroInt, _ := strconv.Atoi(numero)
	if numeroInt > 100 {
		texto = "El numero es mayor a 100"
	} else {
		texto = "El numero es menor a 100"
	}
	return numeroInt, texto
}
