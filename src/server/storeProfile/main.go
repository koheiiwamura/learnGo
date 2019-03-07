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

var Persons []Person

func ProfileAdd(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var storePerson Person
	// 構造体に当てはめる
	var decode_err = json.NewDecoder(r.Body).Decode(&storePerson)
	if decode_err != nil {
		http.Error(w, fmt.Sprintf("%d bad Request", http.StatusBadRequest), http.StatusBadRequest)
	}

	// 過去のPersonと名前が被らないかチェックする
	for i := 0; i < len(Persons); i++ {
		if Persons[i].Name == storePerson.Name {
			http.Error(w, fmt.Sprintf("%d bad Request", http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
	Persons = append(Persons, storePerson)
	// 確認用のログを出力しておく
	log.Printf("\nname: %s\nage: %d\ngender: %s\nfavorite_foods: %s\n", storePerson.Name, storePerson.Age, storePerson.Gender, storePerson.FavoriteFood)
	log.Printf("\nnum of Persons: %d\n", len(Persons))
	// 201を返答
	w.WriteHeader(http.StatusCreated)
}

func main() {
	router := httprouter.New() // HTTPルーターを初期化

	router.POST("/Profile/add", ProfileAdd)

	// Webサーバーを8080ポートで立ち上げる
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
