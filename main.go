package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

var posts = []Article{}
var showPost = Article{}
var env = goDotEnvVariable("ENV")
var AbsolutePath = "/home/avutzhan/go-workspace/src/go-website/templates"

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func index(w http.ResponseWriter, r *http.Request) {
	if env == "production" {
		AbsolutePath = "/var/www/go-website/templates"
	}
	t, err := template.ParseFiles(AbsolutePath+"/index.html", AbsolutePath+"/header.html", AbsolutePath+"/footer.html")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	if env == "production" {
		AbsolutePath = "/var/www/go-website/templates"
	}
	t, err := template.ParseFiles(AbsolutePath+"/create.html", AbsolutePath+"/header.html", AbsolutePath+"/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Not all data filled")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES('%s', '%s', '%s')", title, anons, full_text))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	if env == "production" {
		AbsolutePath = "/var/www/go-website/templates"
	}

	t, err := template.ParseFiles(AbsolutePath+"/show.html", AbsolutePath+"/header.html", AbsolutePath+"/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	showPost = Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			panic(err)
		}

		showPost = post
	}

	t.ExecuteTemplate(w, "show", showPost)

}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	handleFunc()
}
