package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name string
	age uint16
	money int16
	avg_grades, happiness float64

}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Username is %s. He is %d. " +
		"He has a %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"bob", 25, -50, 4.2, 0.8 }
	bob.setNewName("Alehandro") //function work with copy by default so this is another object
	// you must to give a pointer of object to change it so you use one object in both operations
	fmt.Fprintf(w, bob.getAllInfo()) // and this is another object uint16 cant be minus int16 can be minus
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
	//var bob User = ...
	//bob := User{name: "bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8 }

	handleRequest()
}
