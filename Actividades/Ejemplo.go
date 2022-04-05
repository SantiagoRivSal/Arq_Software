package main

import (
	"fmt"
	//"os"
)

//func main() {

	/*sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}*/

	/*
		for i := 0; i < 10; i++ {
			defer fmt.Println(i)
			//defer fmt.Printf("%d", i)
		}
	*/
	/*
		f, _ := os.Create("data.txt") //*os.File
		fmt.Fprintln(f, "data")
		f.Close()
	*/
	/*f, _ := os.ReadFile("data.txt")
	fmt.Print(string(f))*/

//}
func Hello() string{
	return "Hello, world"
}
func main {
	fmt.Println(Hello())
}