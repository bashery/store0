package main

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Users struct {
	//gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phon       string `json:"phon"`
	Avatarlink string `json:"avatarlink"`
}
