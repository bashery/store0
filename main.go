package main

//"fmt"
//"github.com/gin-gonic/gin"
//"github.com/jinzhu/gorm"
//"github.com/jinzhu/gorm/dialects/mysql"

func init() {

	setdb()
}
func main() {
	// set up db and connec
	defer db.Close()

	//toures functions
	router()

}
