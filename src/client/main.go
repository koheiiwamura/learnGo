package main

import (
	"flag"
	"fmt"
	"net/http"
)

type arrayFlags []string

func callGetProfile(name string) {
	url := "http://localhost:8080/Profile/" + name
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	response, _ := client.Do(req)
	defer response.Body.Close()
	fmt.Println(response)
	return
}

func callStoreProfile() {
	fmt.Println("hello")
}

func main() {
	var (
		name   = flag.String("name", "", "name flug")
		age    = flag.Int("age", 0, "age flug")
		gender = flag.String("gender", "", "gender flug")
	)
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	fmt.Println(*name, *age, *gender)
	callGetProfile(*name)
}
