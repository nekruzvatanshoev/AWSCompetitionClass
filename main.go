package main

import (
	"fmt"
	"log"
	"net/http"
)


var globalCounter = 0

func Home(w http.ResponseWriter, r *http.Request) {
	globalCounter +=1

	fmt.Fprintln(w, fmt.Sprintf("Welcome to homepage! You have visited homepage %d times!",globalCounter))
	log.Println(fmt.Sprintf("Welcome to homepage! You have visited homepage %d times!",globalCounter))
}



func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Home)
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatal(err)
	}
}

