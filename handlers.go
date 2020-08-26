package main

import "github.com/gin-gonic/gin"

func index(c *gin.Context) {
	c.HTML(200, "test.html", gin.H{
		"title": "Test website",
	})
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
