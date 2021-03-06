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
	r.LoadHTMLGlob("templates/*html")

	//r.GET("/incr", )
	//r.GET("/logout", logout)

	r.GET("/sign", signup)
	r.POST("/newuser", newUser)
	r.POST("/login", login)
	r.GET("/user", getUser)

	r.GET("/", home)

	r.GET("/acount/:name", acount)
	r.GET("/login", loginPage)
	r.GET("/stores", stores)
	r.GET("/mystore", mystore)
	//r.GET("/about", about)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
