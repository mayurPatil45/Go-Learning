package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
	// log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to golang series </h1> <h3> Hi </h3> <button>Click Me</button>"))
}

// go build . := generates exe file
// go mod tidy := 