package router

import (
	"api/src/controller"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/users", controller.Create)
	r.GET("/users", controller.List)
	r.GET("/users/:id", controller.GetByID)
	r.PUT("/users/:id", controller.UpdateByID)
	r.DELETE("/users/:id", controller.DeleteByID)

	return r
}

func AddPostRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/posts", nil)
	r.GET("/posts", nil)
	r.GET("/posts/:id", nil)
	r.PUT("/posts/:id", nil)
	r.DELETE("/posts/:id", nil)

	return r
}
