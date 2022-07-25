package http

import "errors"

type Error struct {
	Error string
}

var (
	ErrEmptyBody = errors.New("empty request body")
)
