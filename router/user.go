package router

import (
	"github.com/gin-gonic/gin"
)

type users struct {
	Name string `json:"name" uri:"name" form:"name"`
	// age  int    `json: "age" uri: "age" form:"age"`
}

func GetUsers(c *gin.Context) {
	var user users

	c.Bind(&user)

	c.JSON(200, gin.H{
		"result": user,
	})
}
