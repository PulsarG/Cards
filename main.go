package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	m, err := template.ParseFiles("html/main.html", "html/header.html", "html/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	m.ExecuteTemplate(w, "main", nil)
}

func List (w http.ResponseWriter, r *http.Request) {
	m, err := template.ParseFiles("html/list.html", "html/header.html", "html/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	m.ExecuteTemplate(w, "list", nil)
}


// *****************************************************************************************************************************************************************

func StartFunc() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))

	http.HandleFunc("/", MainPage)
	http.HandleFunc("/list", List)

	http.ListenAndServe(":5500", nil)
}

func main() {
	StartFunc()
}
