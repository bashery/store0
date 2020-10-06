package main

func init() {
	setdb()
}
func main() {
	// set up db and connec
	defer db.Close()

	//toures functions
	router()

}
