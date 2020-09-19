package users

import (
	"api/models/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetOne(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var user users.User

	if err, message := users.GetOne(_id, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": message})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func List(c *gin.Context) {
	var userList []users.User

	if err, message := users.List(&userList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userList})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parametros enviados invalidos"})
		return
	}

	if err, message := users.Update(_id, user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	user.Id = _id

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	err, message := users.Delete(_id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": id}})
}

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parametros enviados invalidos"})
		return
	}

	if err, message := users.New(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
