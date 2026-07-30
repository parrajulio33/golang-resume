package repositories

import (
	"errors"
	"fmt"

	"github.com/supabase-community/supabase-go"

	"resume-app/models"
)

func Login(client *supabase.Client, email, password string) (models.User, error) {
	user, err := client.Auth.SignInWithEmailPassword(email, password)

	fmt.Print("User: ", user, "\n")

	if err != nil {
		return models.User{}, err
	}

	if user == nil {
		return models.User{}, errors.New("user not found")
	}

	return models.User{}, nil
}
