package main

import (
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func callback() {
	fmt.Println("Callback")
}
func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting the server...")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, nil
}
