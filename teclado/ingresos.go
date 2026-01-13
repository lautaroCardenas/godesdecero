package teclado

import (
	"bufio" //paquete para leer desde teclado
	"fmt"
	"os"      //paquete para manejar entradas y salidas
	"strconv" //paquete para conversiones
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin) //creo un scanner que lee desde la entrada estandar

	fmt.Println("Ingrese un numero 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text()) //convierte el texto ingresado a entero
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error()) //panic detiene la ejecucion del programa
		}
	}

	fmt.Println("Ingrese un numero 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text()) //convierte el texto ingresado a entero
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)

}
