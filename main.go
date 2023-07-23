package main

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

func main() {

	print()

}

func print() {

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

func format_line(data string, size int) {

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

func insert_Struct(User request) {
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
