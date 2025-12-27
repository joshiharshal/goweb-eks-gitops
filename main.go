package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func servicesPage(w http.ResponseWriter, r *http.Request) { // renamed correctly
	http.ServeFile(w, r, "static/services.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("/home", homePage)           // Home page at /
	http.HandleFunc("/services", servicesPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
