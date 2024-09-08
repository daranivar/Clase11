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
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, Users)
		return
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		idRecieved := c.Param("id")
		idRecievedInt, err := strconv.Atoi(idRecieved)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		for k, u := range Users {
			if u.ID == idRecievedInt {
				Users = append(Users[:k], Users[k+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("User %d deleted", idRecievedInt)})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user %d not found", idRecievedInt)})

	})

	r.Run(":8080")
}

// Desde postman utilizo el verbo delete a la url http://localhost:8080/users/3
