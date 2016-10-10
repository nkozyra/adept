package adept

import (
	_"encoding/json"
	)"io/ioutil"
)

var (
	Conf Config
)

type Config struct {
	DB ConfigDB `json:"database"`
}

type ConfigDB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadConfig(f string) {
	fmt.Println("Doing")
}
