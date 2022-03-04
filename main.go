package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Quote struct {
	Quote     string
	Character string
}

func getQuote() Quote {
	response, err := http.Get("https://animechan.vercel.app/api/random")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Quote
	json.Unmarshal(responseData, &responseObject)

	var quote Quote = Quote{Quote: responseObject.Quote, Character: responseObject.Character}

	return quote

}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("main.html")
	var quote = getQuote()
	err := t.Execute(w, quote)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
