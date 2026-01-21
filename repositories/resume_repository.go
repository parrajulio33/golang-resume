package repositories

import (
	"errors"

	"github.com/supabase-community/supabase-go"

	"resume-app/models"
)

func GetResume(client *supabase.Client) (models.Resume, error) {
	var resumes []models.Resume

	_, err := client.
		From("resumes").
		Select("*", "", false).
		Limit(1, "").
		ExecuteTo(&resumes)

	if err != nil {
		return models.Resume{}, err
	}

	if len(resumes) == 0 {
		return models.Resume{}, errors.New("resume not found")
	}

	return resumes[0], nil
}
