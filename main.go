package main

import (
	"log"
	"net/http"

	"golearn/di"
)

func main() {
	// di.Greet(os.Stdout, "Chad")
	log.Fatal(
		http.ListenAndServe("127.0.0.1:8000",
			http.HandlerFunc(di.MyGreeterHandler)))
}
