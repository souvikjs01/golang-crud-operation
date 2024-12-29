package routes

import (
	"crud-operation/controllers"

	"github.com/gin-gonic/gin"
)

func CrudRouters(r *gin.Engine) {
	r.POST("/add-employee", controllers.AddEmployee)
	r.GET("/users", controllers.GetEmployee)
	r.GET("/user/:id", controllers.GetUser)
	r.PUT("/update-user/:id", controllers.UpdateEmployee)
	r.DELETE("/delete-user/:id", controllers.DeleteUser)
}
