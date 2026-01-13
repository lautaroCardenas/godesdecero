package iteraciones

import (
	"fmt"
)

//se puede hacer de otra forma tambien
/*func Interar() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}*/

func Iterar() {
	for j := 100; j > 10; j -= 5 {
		fmt.Println(j)
	}
}
