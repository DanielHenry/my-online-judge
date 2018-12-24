package main

import (
	"net/http"

    "github.com/DanielHenry/my-online-judge/app/Http/Controllers"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setup() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {});
	r.GET("/", func(c *gin.Context) {});
	r.GET("/user/:name", func(c *gin.Context) {});
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	authorized.POST("/", func(c *gin.Context) {});
	authorized.POST("admin", func(c *gin.Context) {});


	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//make login page
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login/index.tmpl", gin.H{
			"title": "Login Page",
		})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	authorized.POST("/", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string);
		c.JSON(http.StatusOK, gin.H{"status": "Logged in!"});
	});
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
