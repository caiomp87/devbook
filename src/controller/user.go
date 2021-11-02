package controller

import (
	"api/src/models"
	"api/src/repositories"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "unable to read data from request: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repositories.UserCollection.Create(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create an user in database: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created sucessfully",
	})
}

func List(c *gin.Context) {
	c.String(200, "listando usuários")
}

func GetByID(c *gin.Context) {
	c.String(200, "buscando usuário")
}

func UpdateByID(c *gin.Context) {
	c.String(200, "atualizando usuário")
}

func DeleteByID(c *gin.Context) {
	c.String(200, "removendo usuário")
}
