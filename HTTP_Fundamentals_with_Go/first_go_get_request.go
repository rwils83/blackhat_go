package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	r1, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	fmt.Printf("Status: %d", r1.StatusCode)
	body, err := ioutil.ReadAll(r1.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	ioutil.WriteFile("Results.txt", body, 0644)
	r1.Body.Close()
}