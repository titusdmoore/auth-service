package main

import (
    "fmt"
    "os"
    "net/http"
    "html/template"
    "log"
)

type Page struct {
    Title string
    Body []byte
}

func (p *Page) Save() error {
    filname := "pages/" + p.Title + ".txt"
    return os.WriteFile(filname, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := "pages/" + title + ".txt"
    body, error := os.ReadFile(filename)

    if error != nil {
        return nil, error
    }

    return &Page{Title: title, Body: body}, nil
}

func pageViewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func pageEditHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
}

func main() {
    http.HandleFunc("/view/", pageViewHandler)
    http.HandleFunc("/edit/", pageEditHandler)
    http.HandleFunc("/save/", pageSaveHandler)
    log.Fatal(http.ListenAndServe(":1521", nil))
}
