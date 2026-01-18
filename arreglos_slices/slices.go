package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MuestrosSlices() {
	fmt.Println(tablaS)

	procion := arreglo[3:]   // Slice desde la posicion 3 hasta el final
	porcion2 := arreglo[:6]  // Slice desde el inicio hasta la posicion 6
	porcion3 := arreglo[2:7] // Slice desde la posicion 2 hasta la 7

	fmt.Println(procion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

//capacidad es el espacion en memoria asignado al slice
//len es la cantidad de elementos que tiene el slice
//los slicies son dinamicos y pueden crecer o reducir su tamaño
func Capacidad() {
	//crear un slice vacio con make
	elementos := make([]int, 5, 20)                                      //crea un slice de enteros con 5 elementos y capacidad de 20
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos)) //len y cap son funciones predefinidas para slices

	nums := make([]int, 0, 0) //crea un slice vacio
	for i := 0; i < 100; i++ {
		nums = append(nums, i) //append agrega un elemento al final del slice
	}
	fmt.Printf("\nLargo %d, capacidad %d", len(nums), cap(nums))
}
