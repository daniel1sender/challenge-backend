package entities

import (
	"fmt"

	"github.com/daniel1sender/alura-flix/pkg/domain/validate"
	"github.com/google/uuid"
)

type Video struct {
	ID          string
	Title       string
	Description string
	URL         string
}

func NewVideo(title, description, url string) (Video, error) {

	if err := validate.ValidateTitle(title); err != nil {
		return Video{}, fmt.Errorf("could not validate the title due to: %w", err)
	}

	if err := validate.ValidateDescription(description); err != nil {
		return Video{}, fmt.Errorf("could not validate the description due to: %w", err)
	}

	if err := validate.ValidateURL(url); err != nil {
		return Video{}, fmt.Errorf("could not validate the url due to: %w", err)
	}

	return Video{
		ID:          generateID(),
		Title:       title,
		Description: description,
		URL:         url,
	}, nil
}

func generateID() string {
	return uuid.NewString()
}
