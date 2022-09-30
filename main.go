package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template:%v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template:%v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
}

type User struct {
	Name        string
	Age         int
	MovieGenres []string
	Friends     [3]string
	SomeMap     map[string]int
	Meta        MetaUser
}
type MetaUser struct {
	Visits int
}

func executeTemplateWithUser(w http.ResponseWriter, filepath string, user User) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template:%v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, user)
	if err != nil {
		log.Printf("executing template:%v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "template/home.gohtml")
}

func contact(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "template/contact.gohtml")

}

func faq(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "template/faq.gohtml")
}

func greeting(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:        "Ahmed",
		Age:         29,
		MovieGenres: []string{"Comedy", "Action", "Fantasy", "Historical"},
		Friends:     [3]string{"Anwar", "Ousema", "Zied"},
		SomeMap: map[string]int{
			"test":  1,
			"test2": 6,
			"test3": 12,
		},
		Meta: MetaUser{
			Visits: 3,
		},
	}
	executeTemplateWithUser(w, "template/greeting.gohtml", user)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Get("/", home)
	r.Get("/contact", contact)
	r.Get("/faq", faq)
	r.Get("/greeting", greeting)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port 3000....")
	http.ListenAndServe(":3000", r)
}
