package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Request struct {
	Name        string
	Addres      string
	Age         int
	Nymberphone string
}

type people struct {
	name        string
	addres      string
	age         int
	nymberphone string
}

type Edit_data struct {
	Name        string `json:"name"`
	Addres      string `json:"addres"`
	Age         string `json:"age"`
	NumberPhone string `json:"number"`
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

	result = search_account(login)

	file, _ := os.OpenFile("./templates/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	var logg string

	if result == true {
		logg = fmt.Sprintf("Successful login attempt || Username: %s \n", login.username)
		file.WriteString(logg)
	} else {
		logg = fmt.Sprintf("Unsuccessful login attempt || Username: %s || Password: %s  \n", login.username, login.password)
		file.WriteString(logg)
	}

	defer file.Close()

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
		result = false
	}

	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	sec := []Usersdata{}

	for rows.Next() {
		pas := Usersdata{}
		err := rows.Scan(&pas.username, &pas.password)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sec = append(sec, pas)
	}

	for _, user := range sec {

		if user.username == login.username {
			result = false
			return
		} else {
			continue
		}

	}

	defer db.Close()

	file, _ := os.OpenFile("./templates/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	var logg string

	_, err = db.Exec("insert into users (username, password) VALUES ($1, $2)", login.username, login.password)
	if err != nil {
		logg = fmt.Sprintf("Unsuccessful attempt to create an account")
		file.WriteString(logg)
		result = false
	} else {

		logg = fmt.Sprintf("Successful attempt to create an account || Username: %s \n", login.username)
		file.WriteString(logg)
		result = true
	}

	return

}

func Send_data_web() (peoples []Request) { //получаем и отправляем данные из базы данных на страницу

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

	for rows.Next() {
		p := Request{}
		err := rows.Scan(&p.Name, &p.Addres, &p.Age, &p.Nymberphone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		peoples = append(peoples, p)
	}

	return peoples

}

func Change_data(originalData Edit_data, editedData Edit_data) (result bool) {
	db, err := sql.Open("postgres", "postgres://postgres:108@localhost/peoples?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	originalAge, _ := strconv.Atoi(originalData.Age)
	editedAge, _ := strconv.Atoi(editedData.Age)

	_, err = db.Exec(`
    	UPDATE peoples
    	SET Name = $1, Addres = $2, Age = $3, NumberPhone = $4
    	WHERE (Name, Addres, Age, NumberPhone) = ($5, $6, $7, $8)`,
		editedData.Name, editedData.Addres, editedAge, editedData.NumberPhone,
		originalData.Name, originalData.Addres, originalAge, originalData.NumberPhone,
	)
	if err != nil {
		fmt.Println("Error updating record:", err)
		return false
	}

	return true
}

func insert_Struct_database_for_users(User Request) { //Новая запись в базу данных для ВЫВОДА НА САЙТЕ (ЕШЁ НЕ ИСПОЛЬЗОВАННО)
	db, err := sql.Open("postgres", "postgres://postgres:108@localhost/peoples?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("insert into peoples (name, addres, age, numberphone) VALUES ($1, $2, $3, $4)", User.Name, User.Addres, User.Age, User.Nymberphone)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Запись успешна")
}

func Edit_cells(data Edit_data) {
	fmt.Println(data)
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
	peoples := []Request{}

	for rows.Next() {
		p := Request{}
		err := rows.Scan(&p.Name, &p.Addres, &p.Age, &p.Nymberphone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		peoples = append(peoples, p)
	}

	max := maxlens{}

	for _, p := range peoples {
		if max.maxlen_name < len(p.Name) {
			max.maxlen_name = len(p.Name)
		}
		if max.maxlen_addres < len(p.Addres) {
			max.maxlen_addres = len(p.Addres)
		}
		if max.maxlen_age < len(strconv.Itoa(p.Age)) {
			max.maxlen_age = len(strconv.Itoa(p.Age))
		}
		if max.maxlen_number < len(p.Nymberphone) {
			max.maxlen_number = len(p.Nymberphone)
		}
	}

	format_line("Name", max.maxlen_name)
	format_line("Addres", max.maxlen_addres)
	format_line("Age", max.maxlen_age)
	format_line("Phone", max.maxlen_number)
	fmt.Println()
	for _, p := range peoples {
		format_line(p.Name, max.maxlen_name)
		format_line(p.Addres, max.maxlen_addres)
		format_line(strconv.Itoa(p.Age), max.maxlen_age)
		format_line(p.Nymberphone, max.maxlen_number)
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
