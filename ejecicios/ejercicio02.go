//crear un tabla de multiplcar y que el usuario ingrese el numero por teclado
// de la tabla que quiere. y validar si hay error y volver a pedir el numero

package ejecicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin) //creo un scanner que lee desde la entrada estandar

	for {
		fmt.Println("Ingrese un numero para ver su tabla de multiplicar: ")
		scanner.Scan()
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error al convertir el numero. Ingrese un numero valido.")
			continue // si hay error, vuelve al inicio del for
		}
		break
	}
	fmt.Printf("Tabla de multiplicar del %d:\n", numero)
	for i := 1; i <= 10; i++ {
		resultado := numero * i //el := crea la variable resultado y le asigna el valor
		texto += fmt.Sprintf("%d x %d = %d\n", numero, i, resultado)
	}
	return texto
}
