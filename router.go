package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func router() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// static files //r.StaticFS("/static", http.Dir("static")) //r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.Static("/a", "./assets")

	//r.LoadHTMLFiles("/templates/index.html", "/templates/test.html") //r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/incr", session)
	// api
	r.POST("/newuser", newUser)
	r.POST("/login", authLogin)
	r.GET("/user", getUser)

	// templates
	r.GET("/", home)
	r.GET("/acount", acount)
	r.GET("/login", login)
	r.GET("/sign", sign)
	r.GET("/stores", stores)
	r.GET("/mystore", mystore)
	//r.GET("/about", about)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
