package files

import (
	//capturar datos de archivos o teclado
	"bufio"
	"fmt"
	"os"

	"github.com/lautarocardenas/godesdecero/ejecicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejecicios.TablaMultiplicar()
	archivo, err := os.Create(filename) //crea el archivo
	if err != nil {
		fmt.Print("Error al crear el archivo" + err.Error())
		return //si hay error, sale de la funcion
	}
	fmt.Fprintf(archivo, texto) //escribe en el archivo
	archivo.Close()             //cierra el archivo siempre hay que cerrar el archivo
}

// esto es para poder usarlo desde main.go y que grabe la tabla en un archivo y no se borre al ejecutar otra vez
func Sumatabla() {
	var texto string = ejecicios.TablaMultiplicar()
	if !Append(filename, texto) { //llama a la funcion Append para agregar texto al final del archivo
		fmt.Println("Error al concatenar contenido")
	}
}

// funcion que agrega texto al final de un archivo sin borrar lo que ya tiene
func Append(filen string, texto string) bool {
	//os.O_WRONLY|os.O_APPEND modo escritura y agregar al final y 0644 permisos de archivo
	archivo, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644) //abre el archivo en modo escritura y agrega al final
	if err != nil {
		fmt.Print("Error al abrir el archivo" + err.Error())
		return false
	}
	_, err = archivo.WriteString(texto) //escribe el texto en el archivo
	if err != nil {
		fmt.Print("Error al escribir en el archivo" + err.Error())
		return false
	}
	archivo.Close() //cierra el archivo
	return true
}

// lectura de archivos
func LeoArchivo() {
	archivo, err := os.Open(filename) //abre el archivo en modo lectura
	if err != nil {
		fmt.Print("Error al abrir el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo) //crea un scanner para leer el archivo
	for scanner.Scan() {                 //lee el archivo linea por linea el scan devuelve true si hay mas lineas
		registro := scanner.Text()   //obtiene el texto de la linea actual
		fmt.Println("> " + registro) //imprime la linea actual
	}
	archivo.Close() //cierra el archivo
}
