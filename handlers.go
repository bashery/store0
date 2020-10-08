// controler page hold handler function of all logic opiration
package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func sessionTest(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}

	session.Set("count", count)
	session.Save()
	c.JSON(200, gin.H{"count": count})
}

func home(c *gin.Context) {
	session := sessions.Default(c)
	//v := session.Get("count")
	u := session.Get("suser")
	fmt.Println(u)
	c.HTML(200, "home.html", u) //gin.H{"session": v})
}

func authLogin(c *gin.Context) {
	time.Sleep(time.Second)
	var user User
	var u User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
	}
	db.First(&u, "email = ?", user.Email)
	if u.Email == user.Email && u.Password == user.Password {

		session := sessions.Default(c)
		var suser string
		v := session.Get("suser")
		if v == nil {
			suser = user.Email
		}
		fmt.Println("suser : ", suser)
		session.Set("suser", suser)
		session.Save()

		c.String(200, "ok")
		return
	}
	c.String(200, user.Username+" not autorizy")
}

// create new register new user in database
func newUser(c *gin.Context) {
	//db.AutoMigrate(&User{}) // Migrate the schema
	var users User
	var err error
	if err = c.BindJSON(&users); err != nil {
		fmt.Println("err is : ", err)
		c.String(200, fmt.Sprintf("%s", err))
		return
	}
	db.Create(&users)
	c.String(200, "ok") //gin.H{"code": "ok"})
}

func getUser(c *gin.Context) {
	var user User
	db.First(&user, "userid = ?", 1) // find product with id 1
	//db.First(&user, "code = ?", "L1212") // find product with code l1212
	c.String(200, user.Avatarlink)
}

// handle product
func newProduct(c *gin.Context) {
	db.AutoMigrate(&Product{}) // Migrate the schema
	var product Product
	if err := c.BindJSON(&product); err != nil {
		fmt.Println("err is : ", err)
	}
	db.Create(&product)
	c.JSON(200, "ok") //gin.H{"code": "ok"})
}

func acount(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("suser")
	c.HTML(200, "acount.html", user)
}

func sign(c *gin.Context) {
	c.HTML(200, "sign.html", nil) // gin.H{})
}

func login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func stores(c *gin.Context) {
	c.HTML(200, "stores.html", nil) // gin.H{})
}

func mystore(c *gin.Context) {
	c.HTML(200, "mystore.html", nil) // gin.H{})
}
