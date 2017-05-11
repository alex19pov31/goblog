package helpers

import (
	"os"
	"encoding/json"
)

type Config struct {
	path string
	DBName string	`json:"db_name"`
	DBHost string	`json:"db_host"`
}

func LoadConfig(path string) *Config {
	conf := Config{path: path}
	conf.read()

	return &conf
}

func (conf Config) Get(name string) string {
	return ""
}

func (conf *Config) read() {
	f, err := os.Open(conf.path)
	defer f.Close()

	if err != nil {
		return
	}

	decoder := json.NewDecoder(f)
	decoder.Decode(conf)
}