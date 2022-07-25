package validate

import "errors"

var (
	ErrInvalidDescription = errors.New("the description has exceeded the limit of characters")
	ErrEmptyDescription   = errors.New("empty description informed")
)

func ValidateDescription(description string) error {
	if len(description) > 150 {
		return ErrInvalidDescription
	}
	if len(description) == 0 {
		return ErrEmptyDescription
	}
	return nil
}
