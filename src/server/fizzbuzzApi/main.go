package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// URLに与えられている:numをintに変換する
	num, error := strconv.Atoi(p.ByName("num"))
	// 正の整数ではない場合に 400 BadRequestを返す
	if (error != nil) || (num < 1) {
		http.Error(w, fmt.Sprintf("%d Bad Request", http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// 文字列を定義
	var response string
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			response += fmt.Sprintf("%d: %s\n", i, "FizzBuzz!")
		} else if i%5 == 0 {
			response += fmt.Sprintf("%d: %s\n", i, "Buzz")
		} else if i%3 == 0 {
			response += fmt.Sprintf("%d: %s\n", i, "Fizz")
		} else {
			response += fmt.Sprintf("%d:\n", i)
		}
	}
	// Client に出力する
	fmt.Fprintf(w, response)
}

func main() {
	router := httprouter.New() // HTTPルーターを初期化

	// FizzBuzzにGETリクエストがあったらFizzBuzz関数にハンドルする
	router.GET("/FizzBuzz/:num", FizzBuzz)

	// Webサーバーを8080ポートで立ち上げる
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
