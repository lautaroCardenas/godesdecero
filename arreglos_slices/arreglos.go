package arreglos_slices

import "fmt"

// Tabla es un arreglo de enteros
var tabla [10]int = [10]int{10, 0, 59, 23, 45, 67, 89, 12, 34, 56}
var matriz [20][30]int // matriz de 20 filas y 30 columnas

// MuestroArreglos muestra el contenido del arreglo tabla
func MuestroArreglos() {
	tabla[7] = 32
	tabla[2] = 54

	tabla2 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	//	Mostrar con for range
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)

}
