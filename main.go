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
	r.GET("/user/info", func(ctx *gin.Context) {
		userid := ctx.Query("id")
		usersex := ctx.Query("sex")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"userid":  userid,
			"usersex": usersex,
		})
	})

	r.Run(":8005") // listen and serve on 0.0.0.0:8080
}
