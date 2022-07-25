package validate

import "errors"

var (
	ErrInvalidTitle = errors.New("the title has exceeded the limit of characters")
	ErrEmptyTitle   = errors.New("empty title informed")
)

func ValidateTitle(title string) error {
	if len(title) > 100 {
		return ErrInvalidTitle
	}
	if len(title) == 0 {
		return ErrEmptyTitle
	}
	return nil
}
