package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type request struct {
	name        string
	addres      string
	age         int
	nymberphone string
}

type maxlens struct {
	maxlen_name   int
	maxlen_addres int
	maxlen_age    int
	maxlen_number int
}

type Usersdata struct {
	username string
	password string
}

func Search_account_map(login_data map[string]interface{}) (result bool) {

	var login Usersdata
	login.username = login_data["username"].(string)
	login.password = login_data["password"].(string)
	//fmt.Println(login)

	result = search_account(login)

	return result
}

func search_account(User Usersdata) (result bool) { //Поиск аккаунта в базе данных users
	db, err := sql.Open("postgres", "postgres://admin:108@localhost/user-account?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []Usersdata{}

	for rows.Next() {
		p := Usersdata{}
		err := rows.Scan(&p.username, &p.password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	for i := 0; i < len(users); i++ {
		if users[i].username == User.username && users[i].password == User.password {
			result = true
			break
		} else {
			result = false
		}
	}

	return result

}

func Create_new_users(login_data map[string]interface{}) (result bool) { //Новая запись в базу данных (users) пользователей для авторизации

	var login Usersdata
	login.username = login_data["username"].(string)
	login.password = login_data["password"].(string)

	db, err := sql.Open("postgres", "postgres://admin:108@localhost/user-account?sslmode=disable")
	if err != nil {
		log.Fatal(err)
		result = false
	}

	defer db.Close()

	_, err = db.Exec("insert into users (username, password) VALUES ($1, $2)", login.username, login.password)
	if err != nil {
		log.Fatal(err)
		result = false
	} else {
		result = true
	}

	return

}

func insert_Struct_database_for_users(User request) { //Новая запись в базу данных для ВЫВОДА НА САЙТЕ
	db, err := sql.Open("postgres", "postgres://postgres:108@localhost/peoples?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("insert into peoples (name, addres, age, numberphone) VALUES ($1, $2, $3, $4)", User.name, User.addres, User.age, User.nymberphone)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Запись успешна")
}

func print() { //Вывод в консоль базу данных

	db, err := sql.Open("postgres", "postgres://postgres:108@localhost/peoples?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from peoples")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	peoples := []request{}

	for rows.Next() {
		p := request{}
		err := rows.Scan(&p.name, &p.addres, &p.age, &p.nymberphone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		peoples = append(peoples, p)
	}

	max := maxlens{}

	for _, p := range peoples {
		if max.maxlen_name < len(p.name) {
			max.maxlen_name = len(p.name)
		}
		if max.maxlen_addres < len(p.addres) {
			max.maxlen_addres = len(p.addres)
		}
		if max.maxlen_age < len(strconv.Itoa(p.age)) {
			max.maxlen_age = len(strconv.Itoa(p.age))
		}
		if max.maxlen_number < len(p.nymberphone) {
			max.maxlen_number = len(p.nymberphone)
		}
	}

	format_line("Name", max.maxlen_name)
	format_line("Addres", max.maxlen_addres)
	format_line("Age", max.maxlen_age)
	format_line("Phone", max.maxlen_number)
	fmt.Println()
	for _, p := range peoples {
		format_line(p.name, max.maxlen_name)
		format_line(p.addres, max.maxlen_addres)
		format_line(strconv.Itoa(p.age), max.maxlen_age)
		format_line(p.nymberphone, max.maxlen_number)
		fmt.Println()
	}

}

func format_line(data string, size int) { //Вспомогательная функция для метода выше

	var answer string
	var str []string

	word := strings.Split(data, "")

	for i := 0; i <= (size + 4); i++ {
		if i == 0 || i == (size+4) {
			str = append(str, "|")
		} else if i >= 2 && i <= (len(word)+1) {
			str = append(str, word[i-2])
		} else {
			str = append(str, " ")
		}
	}

	answer = strings.Join(str, "")
	fmt.Print(answer)

}
