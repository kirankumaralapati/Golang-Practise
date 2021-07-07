package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
)

type Article struct{
	Name string `jason:"Name"`
	Description string `json:"Description"`
	Content string `json:"Content"`
}

type Articles []Article

func allarticles(w http.handleRequests, r http.Request){
	allarticles:= Articles{
		Article{Name:"Kiran",Description:"Test Program",Content:"Never Mind"}
	}
	fmt.Println("Endpoint Hit: Allrticles Endpoint Hit")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}