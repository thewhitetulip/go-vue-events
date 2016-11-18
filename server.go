package main

/*

This is a hard fork of https://github.com/chenkie/vue-events-bulletin

While using a Go backend for Vue, we just need to execute the home/index template.
We send the AJAX requests from Vue and update the Vue based front end accordingly.

There is one catch, Go templates use {{ }} and Vue also uses the same, thus, we need to
change the definition of our Vue instance and make Vue use ${ } rather than {{ }} otherwise this code won't work.

We can give anything instead of ${ }, but this is short and simple to understand.

new Vue({
  el: '#events',

  data: {
    event: { title: '', detail: '', date: '' },
    events: []
  },
  delimiters: ['${', '}'] // by default it is ['{{','}}']
 //other code
  }

*/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

// This will be the index.html
var homeTemplate *template.Template

// This will store all the templates
var templates *template.Template

// Events is the struct which defines our events
// **it is very important to have the fields capitalized**
type Events struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Date   string `json:"date"`
}

// this is the JSON we will send to the front end
// this can be added to a database too.
var events = []Events{{Id: 1, Title: "Do a Skype", Detail: "Skype with guys from NASA", Date: "2015-12-12"},
	{Id: 2, Title: "Martian releases", Detail: "Relocate to Mars", Date: "2016-1-1"},
	{Id: 3, Title: "What is Vue?"},
}

func main() {
	// Look for all .html files in the current directory and parse them
	templates, err := template.ParseGlob("*.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// find the template with the name index.html
	homeTemplate = templates.Lookup("index.html")

	r := mux.NewRouter()
	r.HandleFunc("/api/events/", EventsHandler).Methods("GET")
	r.HandleFunc("/api/events/", AddHandler).Methods("POST")
	r.HandleFunc("/api/events/{id}", DeleteHandler).Methods("DELETE")

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.HandleFunc("/", HomeHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("running server on 8080")

	http.ListenAndServe(":8080", r)
}

//EventsHandler is the http end point which will supply the list of events
//to our vue front end; handles the GET http request for /api/events/
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(events)
}

// HomeHandler will be rendering the index.html template, it is written in Vue
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, nil)
}

// DeleteHandler will handle the DELETE http req for /api/events/
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in the delete handler")
	// Fetch the ID from the request and delete the value from the database.
}

// AddHandler handles the http POST request for /api/events/
func AddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.FormValue("title"))
	// Do something else with the data, put it in database or do some magic
}
