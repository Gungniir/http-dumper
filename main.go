package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Printf("ListenPath: %s", os.Getenv("LISTEN_PATH"))

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc(os.Getenv("LISTEN_PATH"), dump)
	log.Fatal(http.ListenAndServe(":80", r))
}

func dump(w http.ResponseWriter, r *http.Request) {
	log.Println("--Got new request--")
	log.Println("HEADERS:")

	for key, value := range r.Header {
		log.Printf("%s: %s", key, strings.Join(value, ", "))
	}

	log.Println("BODY:")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	log.Printf("Content-Length: %d", r.ContentLength)
	log.Printf("Read: %d", len(body))
	log.Println(string(body))

	w.Write([]byte(os.Getenv("ANSWER")))
}
