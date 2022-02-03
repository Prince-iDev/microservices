package main

import (
	"log"
	"net/http"
	"os"
)

func handler(rw http.ResponseWriter, r *http.Request) {

}

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	pp := handlers.newPackage{l}
	http.ListenAndServe(":5050", nil)
}
