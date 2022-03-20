package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func SuccessResult(data interface{}) Result {
	return Result{"0", "ok", data}
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResult(nil))
	})

	r.POST("/user", func(c *gin.Context) {
		var user struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, SuccessResult(user))
	})

	r.GET("/test/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, SuccessResult(name))
	})

	r.Run(":5000")
}
