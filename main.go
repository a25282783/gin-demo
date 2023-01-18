package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 創建服務
	r := gin.Default()
	// 加載頁面
	r.LoadHTMLGlob("views/*")
	// 加載資源
	r.Static("/public", "./public")
	// 請求參數
	r.GET("/index", func(ctx *gin.Context) {
		userid := ctx.Query("id")
		usersex := ctx.Query("sex")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"userid":  userid,
			"usersex": usersex,
		})
	})
	// 綁定參數並返回json
	r.GET("/user/:id/:name", func(ctx *gin.Context) {
		id := ctx.Param("id")
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": name,
		})
	})
	// 表單
	r.POST("/create/post", func(ctx *gin.Context) {
		t := ctx.DefaultPostForm("title", "default title")
		a := ctx.DefaultPostForm("article", "default article")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"t": t,
			"a": a,
		})
	})
	// 404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", "")
	})
	// redirect
	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "http://localhost:8008/index")
	})

	r.Run(":8008") // listen and serve on 0.0.0.0:8080
}
