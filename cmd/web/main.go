package main

import (
	"net/http"

	"fmt"

	"github.com/barzegaranmoein/goModHW/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application action
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application of port %s", portNumber))

	http.ListenAndServe(portNumber, nil)
}
