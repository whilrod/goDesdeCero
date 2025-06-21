package main

import (
	"fmt"

	"github.com/whilrod/goDesdeCero/ejercicios"
)

func main() {
	fmt.Println("Hola mundo")
	/* variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(654767)
	fmt.Println(estado)
	fmt.Println(texto)
	//os := runtime.GOOS
	if os := runtime.GOOS; os == "windows" {
		fmt.Println("Estás en: ", os)
	} else {
		fmt.Println("No es windows")
	}
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Printf("Esto es %s \n", os)
	case "darwin":
		fmt.Printf("Esto es %s \n", os)
	default:
		fmt.Printf("Esto es %s \n", os)
	} */
	numeroInt, texto := ejercicios.DevuelveDosValores("89798987")
	fmt.Println(numeroInt)
	fmt.Println(texto)
}
