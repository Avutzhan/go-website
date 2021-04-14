package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	//db, err := sql.Open("mysql", "root:root@/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//установка данных создание нового пользователя
	//insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Alex ti v poryadke', 25)")
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()

	res, err := db.Query("SELECT `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	//Next если есть строка которую можно еще обработать вернет true если строки ольше нет то false
	for res.Next() {
		var user User
		//Scan существуют ли данные в текущем ряде
		//то есть тут мы вытягиваем данные с каждой строки и добавояем их в обьект который мы создали тут
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User %s with age %d", user.Name, user.Age))
	}

}
