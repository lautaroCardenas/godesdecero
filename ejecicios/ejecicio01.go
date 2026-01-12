package ejecicios

import (
	"strconv"
)

// me devuelve un string y un int
func Ejecicio01(texto string) (int, string) {
	numero, err := strconv.Atoi(texto) //convierte un string a entero
	if err != nil {                    //se llama nil cuando no hay error
		return 0, "Error en la conversion" + err.Error() //concatenacion de strings
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}

}
