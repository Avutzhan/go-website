package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name string
	Age uint16
	Money int16
	Avg_grades, Happiness float64
	Hobbies []string

}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Username is %s. He is %d. " +
		"He has a %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"bob", 25, -50, 4.2, 0.8, []string{"test1", "test2", "test3"} }

	tmpl, err := template.ParseFiles("templates/home_page.html")

	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
