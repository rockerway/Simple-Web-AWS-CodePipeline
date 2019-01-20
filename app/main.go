package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port int = 80
const message string = "Go Language Server ~ dev 0.0.1 review 3"

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.Method + " " + req.URL.String() + " " + req.Proto)
		fmt.Fprintf(res, message)
	})
	http.HandleFunc("/healthcheck", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("ok"))
	})
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Server has started to listen at port: " + strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
