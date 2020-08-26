package main

import (
	_ "fmt"
	//"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// set up db and connec
	setdb()
	defer db.Close()

	//toures functions
	router()
}
