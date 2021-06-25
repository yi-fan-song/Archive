package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewClient() (client *Client, err error) {
	c := Client{}

	db, err := gorm.Open(sqlite.Open("testdb.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	c.db = db

	return &c, nil
}
