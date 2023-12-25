package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aswinayyolath/goapimongoDB/router"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Application is listening on port 4000...")
}
