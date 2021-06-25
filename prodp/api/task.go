package api

import (
	"net/http"
	"prodp/data"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID          uint   `JSON:"id"`
	Title       string `JSON:"title"`
	Subtitle    string `JSON:"subtitle"`
	Description string `JSON:"description"`
	IsComplete  bool   `JSON:"isComplete"`
}

type TaskResponse struct {
	Task
	LastModified time.Time `JSON:"lastModified"`
}

func getTask(c echo.Context, repo *data.Client) error {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	record := repo.FetchTask(uint(taskId))

	return c.JSON(http.StatusOK, TaskResponse{
		Task: Task{
			Title:       record.Title,
			Subtitle:    record.Subtitle,
			Description: record.Description,
			IsComplete:  record.IsComplete,
		},
		LastModified: record.UpdatedAt,
	})
}

func postTask(c echo.Context, repo *data.Client) error {
	var payload Task
	if err := c.Bind(&payload); err != nil {
		return err
	}

	repo.AddTask(data.Task{
		Title:       payload.Title,
		Subtitle:    payload.Subtitle,
		Description: payload.Description,
		IsComplete:  payload.IsComplete,
	})

	return c.NoContent(http.StatusNoContent)
}

func putTask(c echo.Context, repo *data.Client) error {
	var payload Task
	if err := c.Bind(&payload); err != nil {
		return err
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	repo.UpdateTask(uint(taskId), data.Task{
		Title:       payload.Title,
		Subtitle:    payload.Subtitle,
		Description: payload.Description,
		IsComplete:  payload.IsComplete,
	})

	return c.NoContent(http.StatusNoContent)
}
