package main

import (
	"api/routes/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/")
	users.User(v1)

	router.Run("api:3001")
}
