package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// internetten json to go code generatorünü kullanabilirsin.
type Todo struct {
	UserId    int    `json:"userId"` //API'deki verileri eşleştirme işlemi
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) //byte array çeviriyor.

	bodyString := string(bodyBytes) //Byte to string dönüşümü yapıldı.
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo) //gelen json veriyi todo nesnesine çeviriyor.
	fmt.Println(todo)
}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	jsonTodo, _ := json.Marshal(todo) //Elindeki veriyi json formatına çevirdi.

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo)) //Content type dediği gönderilen datanın hangi formatta olduğunu yazacaksın. Birden fazla yazmak için ; ile ayır.
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) //byte array çeviriyor.

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}
