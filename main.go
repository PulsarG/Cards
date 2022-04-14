package main

import (
	"database/sql"
	"fmt"
	"html/template"
	//"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
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

	res, err := database.Query(fmt.Sprintf("SELECT * FROM `words` WHERE freq > 0"))
	if err != nil {
		panic(err)
	}
	defer res.Close()

	wordsarray := []WordsStruct{}

	for res.Next() {
		var sw WordsStruct
		err := res.Scan(&sw.Id, &sw.Fword, &sw.Sword, &sw.Freq)
		if err != nil {
			fmt.Println(err)
			continue
		}
		wordsarray = append(wordsarray, sw)
	}

	show := Next(wordsarray)

	m.ExecuteTemplate(w, "main", show)
	//w.Header().Set("Content-Type", "text/html")
	//m.ExecuteTemplate(w, "main", show)

}


func Next(st []WordsStruct) WordsStruct {
	rand.Seed(time.Now().Unix())
	shw := st[rand.Intn(len(st))]
	return shw
}

func List(w http.ResponseWriter, r *http.Request) {
	m, _ := template.ParseFiles("html/list.html", "html/header.html", "html/footer.html")

	res, err := database.Query(fmt.Sprintf("SELECT * FROM `words`"))
	if err != nil {
		panic(err)
	}
	defer res.Close()

	wordsarray := []WordsStruct{}

	for res.Next() {
		var sw WordsStruct
		err := res.Scan(&sw.Id, &sw.Fword, &sw.Sword, &sw.Freq)
		if err != nil {
			fmt.Println(err)
			continue
		}
		wordsarray = append(wordsarray, sw)
	}

	m.ExecuteTemplate(w, "list", wordsarray)
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
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html/"))))

	rtr.HandleFunc("/", MainPage)
	rtr.HandleFunc("/list", List)
	rtr.HandleFunc("/addwords", AddWords).Methods("POST")
	//rtr.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":5500", nil)
}

func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wordcard")
	if err != nil {
		//log.Println(err)
		panic(err)
	}

	database = db
	defer db.Close()

	StartFunc()
}
