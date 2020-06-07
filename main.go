package main

import (
	"fmt"
	"log"
	"net/http"
	"student_api/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	route := routes.Routes()
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", route))

}
