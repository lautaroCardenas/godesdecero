package variables

import (
	"fmt"
	"strconv" //paquete para conversiones
	"time"    //paquete para manejar fechas y horas que ya tiene go
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Lautaro Cardenas"
	Estado = true
	Sueldo = 90000.50
	Fecha = time.Now()
	//fmt es el paquete que permite imprimir en consola
	fmt.Println("Nombre = ", Nombre)
	fmt.Println("Estado = ", Estado)
	fmt.Println("Sueldo = ", Sueldo)
	fmt.Println("Fecha = ", Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
