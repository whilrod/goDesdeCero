package files

import (
	"fmt"
	"os"

	"github.com/whilrod/goDesdeCero/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.PideNumero()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close() // cerramos el archivo
}

func SumaTabla() {
	var texto string = ejercicios.PideNumero()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar datos")
	}
}

func Append(fileName string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al abrir el WriteString" + err.Error())
		return false
	}

	arch.Close()
	fmt.Println("Se ha grabado correctamente")
	return true
}

func LeoArchivo() {
	archivo, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo))
}
