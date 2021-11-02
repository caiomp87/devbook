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

	if err := models.Prepare(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	users, err := repositories.UserCollection.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to list users in database: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetByID(c *gin.Context) {
	ID := c.Param("id")

	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID cannot be empty",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user, err := repositories.UserCollection.GetByID(ctx, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to find an user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateByID(c *gin.Context) {
	ID := c.Param("id")

	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID cannot be empty",
		})
		return
	}

	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "unable to read data from request: " + err.Error(),
		})
		return
	}

	if err := models.Prepare(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repositories.UserCollection.UpdateByID(ctx, ID, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to update user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}

func DeleteByID(c *gin.Context) {
	ID := c.Param("id")

	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID cannot be empty",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repositories.UserCollection.DeleteByID(ctx, ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to delete user: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
