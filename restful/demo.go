package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// farklı cihaz uygulamaları kullanılan bir format oluşturuyoruz
// soap dedigimiz sistem
// json formatı
// map gibi
// get - set olaylarını yapacaz

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:title`
	Completed bool   `json:completed`
}

func Demo() { //baglantı istegi
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
	//jsondan todo ya dönüştürme
}

func Demo2() {
	todo := Todo{1, 2, "alışveriş yapılacak", false}
	// ekranla eşleştiriliyor
	jsonTodo, err := json.Marshal(todo)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	// post oldugu için veri göndermem lazım
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
	//jsondan todo ya dönüştürme
}
