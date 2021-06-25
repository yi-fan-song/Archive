package api

import (
	"net/http"
	"prodp/data"

	"github.com/labstack/echo/v4"
)

type handler struct {
	path       string
	method     string
	handleFunc func(c echo.Context, repo *data.Client) error
}

type handlers []handler

func getHandlers() handlers {
	var paths handlers

	paths = append(paths, handler{path: "/status", method: http.MethodGet, handleFunc: getStatus})
	paths = append(paths, handler{path: "/tasks", method: http.MethodGet, handleFunc: getTasks})
	paths = append(paths, handler{path: "/task/:id", method: http.MethodGet, handleFunc: getTask})
	paths = append(paths, handler{path: "/task", method: http.MethodPost, handleFunc: postTask})
	paths = append(paths, handler{path: "/task/:id", method: http.MethodPut, handleFunc: putTask})

	return paths
}

func getStatus(c echo.Context, _ *data.Client) error {
	return c.String(http.StatusOK, "Prod+ is online.")
}
