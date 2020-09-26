package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	RegisterRoutes()
	MountServer()
}

func RegisterRoutes() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		ExecuteCalc()
		fmt.Fprintf(res, Message("Code Education Rocks !"))
	})
}

func Message(text string) string {
	return "<h1> " + text + " </h1>"
}

func MountServer() {
	fmt.Printf("Starting Server port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is Running")
}

func ExecuteCalc() {
	var x = 0.00001

	for i := 0; i < 1000000; i++ {
		x = math.Sqrt(x+float64(i))/float64(i) + 1
	}
}
