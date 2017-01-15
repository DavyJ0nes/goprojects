package main

import (
	"fmt"
	"golang.org/x/net/context"
	"goprojects/middleware"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing...")
	w.Write([]byte("Oh Hey!"))
}

func paniker(w http.ResponseWriter, r *http.Request) {
	panic(middleware.ErrInvalidEmail)
}

func withContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// doing type assertion as bar is an interface type
	bar := ctx.Value("foo").(string)
	w.Write([]byte(bar))
}

func main() {
	logger := middleware.CreateLogger("logger")
	http.Handle("/", middleware.Time(logger, hello))
	http.Handle("/panic", middleware.Recover(paniker))
	http.Handle("/context", middleware.PassContext(withContext))
	http.ListenAndServe(":3000", nil)
}
