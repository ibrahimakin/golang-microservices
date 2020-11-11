package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	data, err := ioutil.ReadAll(r.Body)
	//log.Printf("Data %s\n", data)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		// res.WriteHeader(http.StatusBadRequest)
		// res.Write([]byte("Ooops"))
		return
	}

	fmt.Fprintf(rw, "Hello %s", data)
}
