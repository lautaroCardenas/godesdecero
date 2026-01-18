package funciones

import (
	"fmt"
)

// Tabla crea una funcion closure para generar la tabla de multiplicar
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}

}

// LLamarClosure llama a la funcion closure de la tabla de multiplicar
func LLamarClosure() {
	tabladel := 2
	Mitabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(Mitabla())
	}
}
