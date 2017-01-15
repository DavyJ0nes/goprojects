package main

import (
	"fmt"
	"goprojects/cms"
	"net/http"
)

func main() {
	// Main routes
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/page", cms.ServePage)
	http.HandleFunc("/post", cms.ServePost)
	// Resources
	http.HandleFunc("/css/", cms.ServeResource)
	http.HandleFunc("/images/", cms.ServeResource)

	// Start Server
	fmt.Println("...Listening on 3000")
	http.ListenAndServe(":3000", nil)
}
