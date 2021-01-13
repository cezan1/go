package model

import "cmdedu/util"

type Config struct {
	BasePath string `json:"base_path"`
	DatePath string `json:"date_path"`
}

func NewConfigWithFile(name string) (*Config, error) {
	c := &Config{}
	util.ReadJosn(name, c)
	return c, nil
}
