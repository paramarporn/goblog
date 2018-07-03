package main

import (
	"fmt"

	"github.com/paramarporn/goblog/accountservice/service"
)

/*
func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
*/

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8080")
}
