package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
}

type Article struct {
	Title    string
	Author   string
	Duration int
}

type ArticleData struct {
	Title    string
	Articles []Article
}

func db_setup(db *sql.DB) {

}

func registerHandlers(db *sql.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Page{"Titus' Awesome Blog"}
		t, err := template.ParseFiles("views/index.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, page)
	})

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		articles := []Article{
			{"Hello World", "Titus Moore", 60},
			{"This is an awesome blog post", "Titus Moore", 45},
			{"Facts Don't Care About Your Feelings", "Sutit Eroom", 135},
		}
		data := ArticleData{"Articles", articles}
		t, err := template.ParseFiles("views/components/articles.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, data)
	})
}

func main() {
	db, err := sql.Open("sqlite3", "./data/blog.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db_setup(db)
	registerHandlers(db)

	fs := http.FileServer(http.Dir("./views/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Listening on port :1521")
	http.ListenAndServe(":1521", nil)
}
