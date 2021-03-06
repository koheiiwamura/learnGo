package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Person は基本情報
type Person struct {
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Gender       string   `json:"gender"`
	FavoriteFood []string `json:"favorite_foods"`
}

// Persons は生成したPersonを保存しておくmap
var Persons map[string]Person = map[string]Person{}

func getProfileDetail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	responsePerson, notFound := Persons[name]

	if !notFound {
		http.Error(w, fmt.Sprintf("Not found %s", name), http.StatusNotFound)
		return
	}
	jsonBytes, err := json.Marshal(responsePerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(string(string(jsonBytes))))
}

// AddProfile はpersonを新規追加する関数
func AddProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	defer r.Body.Close()

	var storePerson Person
	// 構造体に当てはめる
	var err = json.NewDecoder(r.Body).Decode(&storePerson)
	if err != nil {
		http.Error(w, fmt.Sprintf("%d bad Request", http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// 過去のPersonと名前が被らないかチェックする
	_, alreadyExists := Persons[storePerson.Name]
	if alreadyExists {
		http.Error(w, fmt.Sprintf("%d bad Request", http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	Persons[storePerson.Name] = storePerson
	// 確認用のログを出力しておく
	log.Printf("\nnum of Persons: %d\n", len(Persons))
	// 201を返答
	w.WriteHeader(http.StatusCreated)
	jsonBytes, err := json.Marshal(storePerson)
	if err != nil {
		http.Error(w, fmt.Sprintf("%d Internal Server Error", http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(string(jsonBytes)))
}

func main() {
	router := httprouter.New() // HTTPルーターを初期化

	router.POST("/Profile/add", AddProfile)

	router.GET("/Profile/:name", getProfileDetail)

	// Webサーバーを8080ポートで立ち上げる
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
