package main

import (
	"errors"
	"fmt"
)

func division(a, b float64) (float64, error) {

	if b == 0 {
		return -1, errors.New("MAL")
	}
	return a / b, nil
}

func main() {

	div, err := division(15, 32)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println("Division:", div)

}
