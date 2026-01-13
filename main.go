package main

import (
	"github.com/lautarocardenas/godesdecero/ejecicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(123456)
	fmt.Println(estado)
	fmt.Println(texto) */

	/*if os := runtime.GOOS; os == "Linux" || os == "OsX" {
		fmt.Println("Esto no es Windows ", os) // el os es una variable local
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Estamos en Linux")
	case "Daerwin":
		fmt.Println("Estamos en Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejecicios.Ejecicio01("10")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros()

	iteraciones.Iterar()*/

	ejecicios.TablaMultiplicar()

}
