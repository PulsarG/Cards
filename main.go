package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	//"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type WordsStruct struct {
	Id    int
	Fword string
	Sword string
	Freq  int
	User  string
}

var database *sql.DB

const CODE = 301

var datas []byte

func LoadWords() ([]WordsStruct, WordsStruct, error) {
	res, err := database.Query("SELECT * FROM `words`")
	if err != nil {
		log.Println(err)
	}
	defer res.Close()

	wordsarray := []WordsStruct{}
	allwordsarray := []WordsStruct{}

	for res.Next() {
		var sw WordsStruct
		err := res.Scan(&sw.Id, &sw.Fword, &sw.Sword, &sw.Freq, &sw.User)
		if err != nil {
			log.Println(err)
			continue
		}
		if sw.Freq != 0 {
			wordsarray = append(wordsarray, sw)
		}
		allwordsarray = append(allwordsarray, sw)
	}

	withoutZeroFreq := wordsarray[rand.Intn(len(wordsarray))]

	return allwordsarray, withoutZeroFreq, nil
}

func MainPage(w http.ResponseWriter, r *http.Request) {

	m, _ := template.ParseFiles("html/main.html", "html/header.html", "html/footer.html")

	/* _, words, err := LoadWords()
	if err != nil {
		log.Panicln(err)
	} */

	helloVar := WordsStruct{Fword: "Нажми меня", Sword: "Press Me"}

	m.ExecuteTemplate(w, "main", helloVar)
}

func SetData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	data := make(map[string]interface{})

	_ = json.Unmarshal(body, &data)

	freqqy := data["freqq"]
	iddy := data["idd"]

	s := fmt.Sprintf("%.0f", freqqy)
	fr, _ := strconv.Atoi(s)

	ss := fmt.Sprintf("%.0f", iddy)
	iddd, _ := strconv.Atoi(ss)

	upd, _ := database.Prepare("UPDATE words SET freq = ? WHERE id = ?")
	upd.Exec(fr, iddd)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	
	_, words, err := LoadWords()
	if err != nil {
		log.Panicln(err)
	}

	datas, _ = json.Marshal(words)

	w.Write(datas)
}

func List(w http.ResponseWriter, r *http.Request) {
	m, err := template.ParseFiles("html/list.html", "html/header.html", "html/footer.html")
	if err != nil {
		panic(err)
	}

	wordsarray, _, err := LoadWords()
	if err != nil {
		log.Println(err)
	}

	m.ExecuteTemplate(w, "list", wordsarray)
}

func AddWords(w http.ResponseWriter, r *http.Request) {
	firstWord := r.FormValue("firstword")
	secondWord := r.FormValue("secondword")
	frequens := 3
	user := "admin"

	database.Exec("INSERT INTO `words` (`firstword`, `secondword`, `freq`, `user`) VALUES (?, ?, ?, ?)", firstWord, secondWord, frequens, user)

	http.Redirect(w, r, "/", CODE)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	m, err := template.ParseFiles("html/userpage.html", "html/header.html", "html/footer.html")
	if err != nil {
		log.Println(err)
	}
	m.ExecuteTemplate(w, "userpage", nil)
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
	rtr.HandleFunc("/user", LoginPage)
	rtr.HandleFunc("/get", GetData)
	rtr.HandleFunc("/set", SetData)
	/* rtr.HandleFunc("/reg", RegNewUser)
	rtr.HandleFunc("/login", LoginEnter) */

	http.ListenAndServe(":5500", nil)
}

func main() {

	rand.Seed(time.Now().Unix())

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wordcard")
	if err != nil {
		log.Println(err)
	}

	database = db
	defer db.Close()

	StartFunc()
}
