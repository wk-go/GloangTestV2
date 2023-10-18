package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "It works!")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	})
	http.HandleFunc("/show-request", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]any{
			"HEADER": r.Header,
			"URL":    r.URL,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		var jsonStr bytes.Buffer
		json.Indent(&jsonStr, jsonData, "", "  ")
		fmt.Fprintln(w, jsonStr.String())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
