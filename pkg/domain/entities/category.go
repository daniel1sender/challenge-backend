package entities

import (
	"fmt"

	"github.com/daniel1sender/alura-flix/pkg/domain/validate"
)

type Category struct {
	ID    string
	Title string
	Color string
}

func NewCategory(title, color string) (Category, error) {

	if err := validate.ValidateTitle(title); err != nil {
		return Category{}, fmt.Errorf("could not validate the title due to: %w", err)
	}

	if err := validate.ValidateColor(color); err != nil {
		return Category{}, fmt.Errorf("could not validate the color due to: %w", err)
	}

	return Category{
		ID:    generateID(),
		Title: title,
		Color: color,
	}, nil
}
