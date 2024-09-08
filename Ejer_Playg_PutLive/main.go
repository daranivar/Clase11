package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binging:"required"`
}

var Users = []User{
	{1, "Juan Pablo"},
	{2, "José Perez"},
	{3, "Jaime Pogo"},
	{4, "Janero Pinto"},
}

func main() {
	r := gin.Default()

	r.PUT("/users/:id", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)

		idRecieved := c.Param("id")
		idRecievedInt, err := strconv.Atoi(idRecieved)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for k, u := range Users {
			if u.ID == idRecievedInt {
				user.ID = idRecievedInt
				Users[k] = user
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", idRecievedInt)})
	})

	r.Run(":8080")
}

////Desde postman envio el header content/type application/json y desde el body las claves {"username": "Nahuel"} al id 2
//en el endpoint http://localhost:8080/users/2
