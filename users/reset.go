package users

import (
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
)

var validTokens []string

const passwordForm = `
<h1>Enter your new password</h1>
<form action="/reset" method="POST">
  <input type="hidden" name="email" value="{{ . }}" required>

	<label for="password">Password</label>
	<input type="password" name="password" required>

	<input type="submit" value="Submit">
</form>
`

func SendPasswordResetEmail(email string) error {
	token := string(genRandBytes())
	validTokens = append(validTokens, token)
	resetLink := "http://localhost:3000/reset?user=" + email + "&token=" + token

	username := os.Getenv("GMAIL_USER")
	password := os.Getenv("GMAIL_PASS")
	auth := smtp.PlainAuth("smtp.gmail.com:587", username, password, "smtp.gmail.com")

	return smtp.SendMail("smtp.gmail.com:587", auth, username, []string{email}, []byte("Click here to reset your passsword: "+resetLink))
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var isValid bool
		var index int

		// Get params from query string
		email := r.URL.Query().Get("user")
		token := r.URL.Query().Get("token")

		// Check if Token is Valid
		for i, tok := range validTokens {
			if tok == token {
				isValid = true
				index = i
			}
		}
		if isValid != true {
			http.Error(w, "Token not valid", http.StatusUnauthorized)
			return
		}

		// Delete Token from validTokens array
		validTokens = append(validTokens[:index], "")

		// render reset template
		t, _ := template.New("password").Parse(passwordForm)
		t.Execute(w, email)
		return

	case "POST":
		// Get new password from form value
		password := r.FormValue("password")
		email := r.FormValue("email")
		r.ParseForm()

		// Reset password by overwriting old password
		err := SetPassword(email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Successully reset password"))
	}
}

// genRandBytes generates a 32 byte long string of random bytes
func genRandBytes() []byte {
	b := make([]byte, 24)
	_, err := rand.Read(b)
	if err != nil {
		// throwing panic because not being able to generate Random number
		// is ususally an indication of OS failure
		panic(err)
	}
	return []byte(base64.URLEncoding.EncodeToString(b))
}
