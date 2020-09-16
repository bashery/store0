package main

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	//gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phon       string `json:"phon"`
	Avatarlink string `json:"avatarlink"`
}

// TODO handle this
func getProduct(c *gin.Context) {
	var product Product
	db.First(&product, "productid = ?", 1) // find product with id 1
	//db.First(&user, "code = ?", "L1212") // find product with code l1212
	c.String(200, product.Code)
}
