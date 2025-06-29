package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PideNumero() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string
	for {
		fmt.Println("Ingresa el numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error, intenta nuevamente...")
				continue
			} else {
				break
			}
		}
	}

	fmt.Printf("Tabla del %d \n\n", numero)
	texto = fmt.Sprintf("Tabla del %d \n\n", numero)

	for i := 1; i <= 10; i++ {
		//fmt.Println(numero + "X " + i + "= " + numero*i)
		texto += fmt.Sprintf("%d X %d = %d \n", numero, i, numero*i)
	}
	texto += "\n"
	return texto
}
