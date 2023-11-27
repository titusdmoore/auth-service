package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func registerHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/index.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
	})
}

func main() {
	registerHandlers()
	fs := http.FileServer(http.Dir("./views/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Listening on port :1521")
	http.ListenAndServe(":1521", nil)
}
