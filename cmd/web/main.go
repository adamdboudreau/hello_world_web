package main

import (
	"fmt"
	"net/http"

	"github.com/adamdboudreau/hello-world-web/pkg/handlers"
)

// mac run bunch of go files together "go run *.go"
// windows run bunch of go files together "go run ."
const portNumber = ":3030"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world !<div>ok</div>")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})*/
	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
