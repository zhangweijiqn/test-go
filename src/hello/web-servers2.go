package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type MyHandler map[string]dollars

func (self MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello!")
	case "/list":
		for item, price := range self {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := self[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	handler := MyHandler{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", handler))
}

/*
curl http://localhost:8000/
curl http://localhost:8000/list
curl http://localhost:8000/price?item=shoes
*/
