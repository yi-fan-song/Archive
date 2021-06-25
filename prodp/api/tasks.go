package api

import (
	"net/http"
	"prodp/data"

	"github.com/labstack/echo/v4"
)

func getTasks(c echo.Context, repo *data.Client) error {
	tasksDTO := repo.FetchTasks()

	if len(tasksDTO) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Could not find any tasks.")
	}

	var tasks []TaskResponse
	for _, task := range tasksDTO {
		tasks = append(tasks, TaskResponse{
			Task: Task{
				Title:       task.Title,
				Subtitle:    task.Subtitle,
				Description: task.Description,
			},
			LastModified: task.UpdatedAt,
		})
	}

	if err := c.JSON(http.StatusOK, tasks); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
