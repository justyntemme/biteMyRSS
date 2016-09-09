package web

import (
	"html/template"
	"net/http"
)

//Page exported stuct which holds page information to be rendered
type Page struct {
	Title string
	Body  string
	Type  string
}

//StartServer exported to start web server at beggining on program start
func StartServer(addr string) {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static", fs))
	http.HandleFunc("/", serveHTTP)
	http.ListenAndServe(addr, nil)

}

func serveHTTP(d http.ResponseWriter, req *http.Request) {
	serveTemplate(d, &Page{Title: "biteMyRSS", Body: "", Type: "home"})

}

func serveTemplate(d http.ResponseWriter, page *Page) {
	d.Header().Add("Content Type", "text/html")
	var file string
	if page.Type == "home" {
		file = "home"
	}
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.ExecuteTemplate(d, file, page)
}
