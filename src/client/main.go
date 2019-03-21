package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const host = "http://localhost:8080/Profile/"

type Person struct {
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Gender       string   `json:"gender"`
	FavoriteFood []string `json:"favorite_foods"`
}

func callGetProfile(name string) {
	url := host + name
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
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
	url := host + "add/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}
	client := new(http.Client)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
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
