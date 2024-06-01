package routes

import (
    "myapi/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/users", controllers.FindUsers)
    r.GET("/users/:id", controllers.FindUser)
    r.POST("/users", controllers.CreateUser)
    r.PUT("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)

    return r
}
