package data

func (c *Client) Migrate() error {
	if err := c.db.AutoMigrate(&Task{}); err != nil {
		return err
	}

	return nil
}
