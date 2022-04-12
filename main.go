package main

import (
	"database/sql"
	"fmt"
	"html/template"
	//"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type WordsStruct struct { // выгрузка БД
	Id    int
	Fword string
	Sword string
	Freq  int
}

var database *sql.DB

func MainPage(w http.ResponseWriter, r *http.Request) {

	m, _ := template.ParseFiles("html/main.html", "html/header.html", "html/footer.html")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wordcard")
	if err != nil {
		//log.Println(err)
		panic(err)
	}

	database = db
	defer db.Close()

	res, err := database.Query("SELECT * FROM `words`")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	wfshow := []WordsStruct{}

	for res.Next() {
		var sw WordsStruct
		err := res.Scan(&sw.Id, &sw.Fword, &sw.Sword, &sw.Freq)
		if err != nil {
			fmt.Println(err)
			continue
		}

		wfshow = append(wfshow, sw)
	}
	//w.Header().Set("Content-Type", "text/html")
	//m.Execute(w, wfshow)
	m.ExecuteTemplate(w, "main", wfshow)
}

func List(w http.ResponseWriter, r *http.Request) {
	m, err := template.ParseFiles("html/list.html", "html/header.html", "html/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	m.ExecuteTemplate(w, "list", nil)
}

func AddWords(w http.ResponseWriter, r *http.Request) {
	firstWord := r.FormValue("firstword")
	secondWord := r.FormValue("secondword")
	frequens := 3

	if firstWord == "" || secondWord == "" {
		http.Redirect(w, r, "/", 301)
	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wordcard")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `words` (`firstword`, `secondword`, `freq`) VALUES ('%s', '%s', '%d')", firstWord, secondWord, frequens))
		if err != nil {
			panic(err)
		}
		defer insert.Close()
		http.Redirect(w, r, "/", 301)
	}
}

// *****************************************************************************************************************************************************************

func StartFunc() {
	rtr := mux.NewRouter()

	http.Handle("/", rtr)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))

	rtr.HandleFunc("/", MainPage)
	rtr.HandleFunc("/list", List)
	rtr.HandleFunc("/addwords", AddWords).Methods("POST")

	http.ListenAndServe(":5500", nil)
}

func main() {
	StartFunc()
}
