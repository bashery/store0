// controler page hold handler function of all logic opiration
package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	// TODO data form validator here,
	var user User
	var u User
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	fmt.Println(user.Email, user.Password)
	db.First(&u, "email = ?", user.Email)
	if u.Email == user.Email && u.Password == user.Password {
		session := sessions.Default(c)
		session.Set("suser", u.Email)
		session.Save()
		c.HTML(200, "home.html", session.Get("suser")) //gin.H{"session": v})
		return
	}
	//c.Redirect(http.StatusMovedPermanently, "/login")
	c.HTML(200, "login.html", "not correct") // TODO use jet template
}

func home(c *gin.Context) {
	session := sessions.Default(c)
	//v := session.Get("count")
	u := session.Get("suser")
	fmt.Println("home suser is :", u)
	c.HTML(200, "home.html", u) //gin.H{"session": v})
}

// create new register new user in database
func newUser(c *gin.Context) {
	//db.AutoMigrate(&User{}) // Migrate the schema
	var user User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phon = c.PostForm("phon")
	db.Create(&user)
	fmt.Println(user)
	c.HTML(200, "login.html", nil) //gin.H{"code": "ok"})
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
	name := c.Param("name")
	if user == name {
		c.HTML(200, "acount.html", []string{name, "my acount"})
		return

	}
	c.HTML(200, "acount.html", name)
}

func signup(c *gin.Context) {
	c.HTML(200, "sign.html", nil) // gin.H{})
}

func loginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func stores(c *gin.Context) {
	c.HTML(200, "stores.html", nil) // gin.H{})
}

func mystore(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get("suser")
	c.HTML(200, "mystore.html", name)
}

/* we should use logout ??
func logout(c *gin.Context) {
	session := sessions.Default(c)
	u := session.Get("suser")
	fmt.Println("clear user session", u)
	session.Clear()
	fmt.Println("suser befor clear", session.Get("suser"))

	c.HTML(200, "home.html", nil)
}
*/

/* func sessionTest(c *gin.Context) {
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
} */
