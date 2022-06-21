package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongo/controllers"
	"net/http"
)

func UserRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.POST("/user", controllers.CreateUser())
	r.GET("/user/:userId", controllers.GetUser())
	r.PUT("/user/:userId", controllers.EditUser())
	r.DELETE("/user/:userId", controllers.DeleteUser())
	r.GET("/users", controllers.GetAllUsers())
}
