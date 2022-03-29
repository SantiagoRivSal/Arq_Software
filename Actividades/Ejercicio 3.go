package main

import (
	"fmt"
	"os"
)

func Crear_Archivo(p string) *os.File {
	fmt.Println("Creando")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func Escribir(f *os.File) {
	fmt.Println("Escribiendo")
	fmt.Fprintln(f, "Hola soy de Senegal, Pero Egipto tenia que ir al Mundial")
}
func Cerrar_Archivo(f *os.File) {
	fmt.Println("Cerrando")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%v/n", err)
		os.Exit(1)
	}
}
func main() {
	f := Crear_Archivo("archivo.txt")
	defer Cerrar_Archivo(f)
	Escribir(f)

}
