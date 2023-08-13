package main

import (
	"net/http"

	"example.com/m/database"
	"github.com/gin-gonic/gin"
)

type users_data struct {
	username string
	password string
}

func main() {
	r := gin.Default()

	r.Static("/static", "./static/")

	r.GET("/login", func(c *gin.Context) {
		r.LoadHTMLFiles("./templates/login_page.html")
		c.HTML(http.StatusOK, "login_page.html", gin.H{
			"title": "Login Page",
		},
		)

	})

	// Обработчик POST-запроса из формы
	r.POST("/login", func(c *gin.Context) {

		var data map[string]interface{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Ошибка декодирования JSON"})
			return
		}

		status := database.Search_account_map(data)

		if status == true {
			c.JSON(http.StatusOK, gin.H{"redirect_url": "/mainpage"})
		} else {
			c.JSON(http.StatusOK, gin.H{"bad_data": "wrong data"})
		}

	})

	r.GET("/mainpage", func(c *gin.Context) {
		r.LoadHTMLFiles("./templates/mainpage.html")
		data := database.Send_data_web()
		c.HTML(http.StatusOK, "mainpage.html", gin.H{
			"people": data,
		})
	})

	r.GET("/registration", func(c *gin.Context) {
		r.LoadHTMLFiles("./templates/registration_page.html")
		c.HTML(http.StatusOK, "registration_page.html", nil)
	})

	r.POST("/registration", func(c *gin.Context) {

		var data map[string]interface{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": "Ошибка декодирования JSON"})
			return
		}

		status := database.Create_new_users(data)

		if status == true {
			c.JSON(http.StatusOK, gin.H{"redirect_url": "/login"})
		} else {
			c.JSON(http.StatusOK, gin.H{"bad_data": "wrong data"})
		}

	})

	r.Run(":8080")
}
