package data

import "errors"

type Config struct {
	Table TableConfig
}

type TableConfig struct {
	Host string
	User string
	Db   string
	Pass string
	Port int
}

func (c *TableConfig) Valid() error {
	if c.Host == "" || c.User == "" || c.Db == "" ||
		c.Pass == "" || c.Port == 0 {
		return errors.New("Unable to create resty configuration.")
	}
	return nil
}
