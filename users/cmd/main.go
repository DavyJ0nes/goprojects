package main

import (
	"fmt"
	"goprojects/users"
	"net/http"
	"os"
)

func main() {
	username, password := os.Getenv("GMAIL_USER"), os.Getenv("GMAIL_PASS")

	err := users.NewUser(username, password)
	if err != nil {
		fmt.Printf("Couldn't create user: %s\n", err.Error())
		return
	}

	err = users.AuthenticateUser(username, password)
	if err != nil {
		fmt.Printf("Couldn't authenticate user: %s\n", err.Error())
		return
	}

	fmt.Println("Successfully created and authenticated user %s", username)

	// reset email
	err = users.SendPasswordResetEmail(username)
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/reset", users.ResetPassword)
	http.ListenAndServe(":3000", nil)
}
