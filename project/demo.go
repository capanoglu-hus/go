package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"prdouctName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProdutc() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) //gelen cevabı okuyacak

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (Product, error) {
	product := Product{ProductName: "ccdd", CategoryId: 1, UnitPrice: 1999.9}
	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		return Product{}, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	return productResponse, nil

}
