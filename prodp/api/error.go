package api

import (
	"fmt"
	"net/http"
)

type Error struct {
	message string
	code    int
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

var NotFoundError = Error{
	code:    http.StatusNotFound,
	message: "Could not find the requested resource",
}
