package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Person struct {
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Gender       string   `json:"gender"`
	FavoriteFood []string `json:"favorite_foods"`
}

var Bob = Person{
	Name:         "Bob",
	Age:          25,
	Gender:       "Man",
	FavoriteFood: []string{"Hamburger", "Cookie", "Chocolate"},
}

var Alice = Person{
	Name:         "Alice",
	Age:          24,
	Gender:       "Woman",
	FavoriteFood: []string{"Apple", "Orange", "Melon"},
}

func Profile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	name := p.ByName("name")
	var responsePerson Person
	switch name {
	case "Bob":
		responsePerson = Bob
	case "Alice":
		responsePerson = Alice
	default:
		http.Error(w, fmt.Sprintf("%d bad Request", http.StatusBadRequest), http.StatusBadRequest)
	}

	jsonBytes, _ := json.Marshal(responsePerson)
	jsonStr := string(jsonBytes)
	fmt.Fprintf(w, jsonStr)
}

func main() {
	router := httprouter.New() // HTTPルーターを初期化

	router.GET("/Profile/:name", Profile)

	// Webサーバーを8080ポートで立ち上げる
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
