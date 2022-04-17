package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Package struct {
	l *log.Logger
}

func NewPackage(l *log.Logger) *Package {
	return &Package{l}
}
func (p *Package) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Waan")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oh an error message", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Oh Oh!!"))
		return
	}

	// log.Printf("Data is:%s\n", d)
	fmt.Fprintf(rw, "\nWhat's on your mind %s?\n", d)
}
