package data

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Subtitle    string
	Description string
	IsComplete  bool
}

type Tasks []Task

func (c *Client) FetchTasks() Tasks {
	var tasks Tasks

	c.db.Find(&tasks)

	return tasks
}

func (c *Client) FetchTask(id uint) Task {
	var task Task

	c.db.First(&task, id)

	return task
}

func (c *Client) AddTask(task Task) {
	c.db.Create(&task)
}

func (c *Client) UpdateTask(id uint, modified Task) {
	var task Task

	c.db.First(&task, id).Updates(modified)
}
