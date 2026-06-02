package main

import (
	"fmt"

	"example.com/go-jwt/controllers"
	"example.com/go-jwt/initializers"
	"example.com/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

// APIを定義。Validateメソッドを呼び出す前にRequireAuthを呼び出して認証を行わせる
func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
	fmt.Println("Hello, World!")
}

