package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template
var templates *template.Template

//Events is the struct which defines our events
type Events struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Date   string `json:"date"`
}

var events = []Events{{Id: 1, Title: "Do a Skype", Detail: "Skype with guys from NASA", Date: "2015-12-12"},
	{Id: 2, Title: "Martian releases", Detail: "Relocate to Mars", Date: "2016-1-1"},
	{Id: 3, Title: "What is Vue?"},
}

func main() {
	templates, err := template.ParseGlob("*.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	homeTemplate = templates.Lookup("index.html")

	r := mux.NewRouter()
	r.HandleFunc("/api/events", EventsHandler).Methods("GET")
	// r.Handle("/public/", http.FileServer(http.Dir("public")))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.HandleFunc("/", HomeHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("running server on 8080")

	http.ListenAndServe(":8080", r)
}

//EventsHandler is the http end point which will supply the list of events
//to our vue front end
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sending JSON")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// d, err := json.Marshal(events)

	// if err != nil {
	// 	fmt.Println("error")
	// } else {
	// 	w.Write(d)
	// }

	json.NewEncoder(w).Encode(events)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, nil)
}
