package main

import "github.com/gin-gonic/gin"

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
	c.HTML(200, "acount.html", gin.H{})
}

func imgs(c *gin.Context) {
	c.HTML(200, "imgs.html", gin.H{})
}

func sign(c *gin.Context) {
	c.HTML(200, "sign.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func stores(c *gin.Context) {
	c.HTML(200, "stores.html", gin.H{})
}

func mystore(c *gin.Context) {
	c.HTML(200, "mystore.html", gin.H{})
}

func user(c *gin.Context) {
	var user Users
	db.First(&user, "userid = ?", 1) // find product with id 1
	//db.First(&user, "code = ?", "L1212") // find product with code l1212
	c.String(200, user.Avatarlink)

}

func newUser(c *gin.Context) {
	// Migrate the schema
	db.AutoMigrate(&Users{})

	// Create
	u := db.Create(&Users{Username: "it is my", Password: "123456", Email: "ab@gorm.com", Phon: "05344667788", Avatarlink: "www.image.com/myimge2.png"})

	c.JSON(200, gin.H{
		"users": u.Value,
	})
}
