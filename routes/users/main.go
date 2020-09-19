package users

import (
	controller "api/controllers/users"
	"github.com/gin-gonic/gin"
)

func User(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/:id", controller.GetOne)
	users.GET("/", controller.List)
	users.DELETE("/:id", controller.Delete)
	users.PUT("/:id", controller.Update)
	users.POST("/", controller.Create)
}
