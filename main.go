package main

import (
	"fmt"

	"github.com/lautarocardenas/godesdecero/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(123456)
	fmt.Println(estado)
	fmt.Println(texto)
}
