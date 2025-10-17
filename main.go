package main

import (
	"crud-operation/config"
	"crud-operation/models"
	"crud-operation/routes"
	"crud-operation/store"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	port := config.GetEnv("PORT")
	dbURL := config.GetEnv("DATABASE_URL")

	store.ConnectDB(dbURL)
	store.DB.AutoMigrate(&models.Employee{})

	fmt.Println("database is successfully connected...")

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to golang crud application",
		})
	})
	routes.CrudRouters(r)

	r.Run(":" + port)
}
