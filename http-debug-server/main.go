package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var Port = flag.Int("port", 1337, "server port")

func main() {

	fmt.Printf("Started on port %v\n\n", *Port)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method + " " + req.URL.String() + " " + req.Proto)
		fmt.Println("Host: " + req.Host)
		for key, value := range req.Header {
			fmt.Printf("%v: %v\n", key, value)
		}
		fmt.Printf("\n")
		body, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(body))
		}
		fmt.Printf("\n\n")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", *Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
