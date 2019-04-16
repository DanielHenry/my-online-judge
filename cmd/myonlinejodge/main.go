package main

import (
	//"fmt"
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/DanielHenry/my-online-judge/internal/pkg/controller/questionscontroller"
	"github.com/DanielHenry/my-online-judge/internal/pkg/controller/userscontroller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/", userscontroller.create)
		v1.GET("/", userscontroller.index)
		v1.GET("/:id", userscontroller.show)
		v1.PUT("/:id", userscontroller.update)
		v1.DELETE("/:id", userscontroller.destroy)
	}
	router.Run("127.0.0.1:8090")
}
