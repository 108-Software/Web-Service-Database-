package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static/")

	r.LoadHTMLFiles("./templates/login_page.html")

	r.GET("/", func(c *gin.Context) {

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

		fmt.Println("Структура POST-запроса:", data)

		c.JSON(200, gin.H{"message": "POST-запрос успешно обработан"})
	})

	r.Run(":8080")
}
