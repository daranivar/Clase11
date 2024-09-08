package main

import "github.com/gin-gonic/gin"

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binging:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var login Login
		c.BindJSON(&login)
		c.JSON(200, gin.H{"status": login.Email})
	})
	r.Run(":8080")
}

//Desde postman envio el header content/type application/json y desde el body lac claves {"email": "Paul Logan","password": "secret"}
