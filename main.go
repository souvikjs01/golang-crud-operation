package main

import (
	"crud-operation/config"
	"crud-operation/models"
	"crud-operation/routes"
	"crud-operation/store"
	"fmt"

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
	routes.CrudRouters(r)

	r.Run(":" + port)
}
