package validate

import (
	"errors"
)

var (
	ErrEmptyURL = errors.New("empty url informed")
)

func ValidateURL(url string) error {
	if url == "" {
		return ErrEmptyURL
	}
	return nil
}
