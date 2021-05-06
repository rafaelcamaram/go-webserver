package main

import (
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	fmt.Printf("Webserver running | PORT: 3000")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
