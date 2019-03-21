package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Person struct {
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Gender       string   `json:"gender"`
	FavoriteFood []string `json:"favorite_foods"`
}

func callGetProfile(name string) {
	url := "http://localhost:8080/Profile/" + name
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	response, _ := client.Do(req)
	defer response.Body.Close()
	fmt.Println(response)
	return
}

func callStoreProfile(name string, age int, gender string, favoriteFoods string) {

	profile := Person{
		Name:         name,
		Age:          age,
		Gender:       gender,
		FavoriteFood: strings.Split(favoriteFoods, " "),
	}

	jsonStr, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}
	url := "http://localhost:8080/Profile/add/"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	client := new(http.Client)
	response, _ := client.Do(req)
	defer response.Body.Close()
	fmt.Println(response)
	return
}

func main() {
	var (
		name          = flag.String("name", "", "name flug")
		age           = flag.Int("age", 0, "age flug")
		gender        = flag.String("gender", "", "gender flug")
		favoriteFoods = flag.String("favorites_foods", "", "favorites_foods")
	)
	flag.Parse()
	if *name != "" && *age == 0 && *gender == "" && *favoriteFoods == "" {
		// 関数に値を渡す前にコピーされてしまうため、ポインタの中身を渡す
		callGetProfile(*name)
	} else if *name != "" {
		// 関数に値を渡す前にコピーされてしまうため、ポインタの中身を渡す
		callStoreProfile(*name, *age, *gender, *favoriteFoods)
	} else {
		fmt.Println("the following arguments are required: name")
	}
}
