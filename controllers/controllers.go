package controllers

import (
	"crud-operation/models"
	"crud-operation/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEmployee(c *gin.Context) {
	var emp models.Employee

	err := c.ShouldBindJSON(&emp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	store.DB.Create(&emp)
	c.JSON(http.StatusCreated, emp)
}

func GetEmployee(c *gin.Context) {
	var emps []models.Employee
	store.DB.Find(&emps)
	c.JSON(http.StatusOK, emps)
}

func GetUser(c *gin.Context) {
	var emp models.Employee
	id := c.Param("id")
	err := store.DB.First(&emp, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func UpdateEmployee(c *gin.Context) {
	var emp models.Employee
	id := c.Param("id")
	err := store.DB.First(&emp, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User Not Found!",
		})
		return
	}
	err2 := c.ShouldBindJSON(&emp)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err2.Error(),
		})
		return
	}
	store.DB.Save(&emp)
	c.JSON(http.StatusOK, emp)
}

func DeleteUser(c *gin.Context) {
	var emp models.Employee
	id := c.Param("id")

	err := store.DB.First(&emp, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found " + err.Error(),
		})
		return
	}
	store.DB.Delete(&emp)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted user",
	})
}
