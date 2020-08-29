package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	//"net/http"
)

// create new register new user ub database
func newUser(c *gin.Context) {
	db.AutoMigrate(&User{}) // Migrate the schema

	var users User
	if err := c.BindJSON(&users); err != nil {
		fmt.Println("err is : ", err)
	}
	db.Create(&users)
	c.JSON(200, "ok") //gin.H{"code": "ok"})
}

func getUser(c *gin.Context) {
	var user User
	db.First(&user, "userid = ?", 1) // find product with id 1
	//db.First(&user, "code = ?", "L1212") // find product with code l1212
	c.String(200, user.Avatarlink)

}

func authLogin(c *gin.Context) {
	var user User
	var u User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
	}

	db.First(&u, "username = ?", user.Username)
	fmt.Println(u)
	if u.Username == user.Username && u.Password == user.Password {
		//c.Redirect(http.StatusFound, "/index")
		c.JSON(200, "ok")
		return
	}
	c.JSON(200, user.Username+" not autorizy")
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Test website",
	})
}

func home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"title": "Test website",
	})
}

func acount(c *gin.Context) {
	c.HTML(200, "acount.html", nil)
}

func imgs(c *gin.Context) {
	c.HTML(200, "imgs.html", gin.H{})
}

func sign(c *gin.Context) {
	c.HTML(200, "sign.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func stores(c *gin.Context) {
	c.HTML(200, "stores.html", gin.H{})
}

func mystore(c *gin.Context) {
	c.HTML(200, "mystore.html", gin.H{})
}
