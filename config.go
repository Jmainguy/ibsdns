package main

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

func config() (credents Config) {
	config_file, err := ioutil.ReadFile("/opt/ibsdns/config.yaml")
	check(err)
	yaml.Unmarshal(config_file, &credents)
	return
}

type Config struct {
	ApiKey   string `json:"apiKey"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}
