package main

import (
	"net/http"
    "fmt"
	"encoding/json"
	"errors"
	"io/util"
)

type Categories [] Category{

}

type Category struct{
	ID string `json:"id"`
	Name string `json:"name"`
}

func main() {
	cats, err := getCategories("MLA")
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println("CATEGORIAS: ", cats)
}

func getCategories(siteID string) (Categories, error) {
	resp, err := http.Get("https://api.mercadolibre.com/sites/MLA/search?q=Motorola") //completar

	bytes := ioutil.ReadAll(resp.Bytes()) //completar

	var cats Categories

	err := json.Unmarshal(bytes, &cats)

	return cats, nil
}