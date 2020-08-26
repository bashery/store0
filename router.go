package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
)

func router() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	// static files
	r.Static("/a", "./assets")
	//r.StaticFS("/static", http.Dir("static"))
	//r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// loading templates
	//r.LoadHTMLFiles("templates/*")
	//r.LoadHTMLFiles("/templates/index.html", "/templates/test.html")
	r.LoadHTMLGlob("templates/*")

	//r.LoadHTMLGlob("templates/*")

	r.POST("/user", newUser)
	r.GET("/user", user)

	// templates
	r.GET("/index", index)
	r.GET("/", home)

	r.GET("/acount", acount)
	r.GET("/login", login)
	r.GET("/sign", sign)
	r.GET("/stores", stores)
	r.GET("/mystore", mystore)
	//r.GET("/about", about)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
