package initializers

import (
	"fmt"
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func InitGoth() {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	if clientID == "" || clientSecret == "" || callbackURL == "" {
		fmt.Println("Google OAuth Credentials are not set")
		return
	}

	goth.UseProviders(
		google.New(clientID, clientSecret, callbackURL, "email", "profile"),
	)
}
