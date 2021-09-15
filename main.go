package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	request1()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := []whoami{
		{
			Name:  "Pavan Gujar",
			Title: "Lead",
			State: "Karnataka",
		},
	}
	json.NewEncoder(response).Encode(who)
	fmt.Println("EndPoint Hit", who)
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to go web API")
	fmt.Println("EndPoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "PavanGujar"
	fmt.Fprintf(response, "About Pavan")
	fmt.Println("EndPoint Hit: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
