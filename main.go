package main

import (
	"fmt"
	"net/http"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/contact":
		contact(w, r)
	case "/faq":
		faq(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>contact me here</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>Q: Is there a free version?</p><p>A: Yes</p><p>Q: Is there a free version?</p><p>A: Yes</p><p>Q: Is there a free version?</p><p>A: Yes</p><p>Q: Is there a free version?</p><p>A: Yes</p>")
}

func main() {
	var router http.HandlerFunc = pathHandler
	fmt.Println("Starting the server on port 3000....")
	http.ListenAndServe(":3000", router)
}
