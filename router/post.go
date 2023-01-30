package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	c.String(http.StatusOK, "posts")
}
