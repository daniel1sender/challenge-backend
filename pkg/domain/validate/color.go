package validate

import "errors"

var (
	ErrInvalidColor = errors.New("the color has exceeded the limit of characters")
	ErrEmptyColor   = errors.New("empty color informed")
)

func ValidateColor(color string) error {
	if len(color) > 100 {
		return ErrInvalidColor
	}
	if len(color) == 0 {
		return ErrEmptyColor
	}
	return nil
}
