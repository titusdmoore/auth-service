package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func notFound(w http.ResponseWriter) {
	renderTemplate(w, "templates/not-found.html", nil)
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
	p, err := loadPage(title)

	if err != nil {
		notFound(w)
	}

	renderTemplate(w, "templates/view.html", p)
}

func pageEditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "templates/edit.html", p)
}

func pageSaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func pageCreateViewHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/create.html", nil)
}

func pageCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "Testing")
	if r.Method != "POST" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
	}

	// fmt.Println(r.GetBody())
	title, body := r.FormValue("title"), r.FormValue("body")
	//
	p := &Page{Title: title, Body: []byte(body)}
	p.Save()
	//
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", pageViewHandler)
	http.HandleFunc("/edit/", pageEditHandler)
	http.HandleFunc("/create/", pageCreateViewHandler)
	http.HandleFunc("/create-post", pageCreateHandler)
	http.HandleFunc("/save/", pageSaveHandler)
	log.Fatal(http.ListenAndServe(":1521", nil))
}
