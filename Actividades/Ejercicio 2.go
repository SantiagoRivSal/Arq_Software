package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Categories []Category

type Category struct {
	ID   string `json:"id"`
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
