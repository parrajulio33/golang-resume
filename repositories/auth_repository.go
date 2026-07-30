package repositories

import (
	"errors"
	"fmt"

	// "fmt"

	"github.com/supabase-community/supabase-go"

	"resume-app/models"
)

func Login(client *supabase.Client, email, password string) (models.User, error) {
	session, err := client.Auth.SignInWithEmailPassword(email, password)

	// fmt.Print("User: ", user, "\n")

	if err != nil {
		return models.User{}, err
	}

	if session == nil {
		return models.User{}, errors.New("user not found")
	}

	displayName, _ := session.User.UserMetadata["display_name"].(string)

	fmt.Print("Display Name: ", displayName, "\n")
	fmt.Print("Email: ", session.User.Email, "\n")
	fmt.Print("Role: ", session.User.AppMetadata["role"], "\n")

	return models.User{
		ID:          session.User.ID,
		Email:       session.User.Email,
		DisplayName: displayName,
	}, nil
}
