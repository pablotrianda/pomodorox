package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "/.config/pomodorox/config.yaml"

type Config struct {
	Timers struct {
		WorkLoops     int `yaml:"workLoops"`
		TimeWorkLoops int `yaml:"timeWorkLoops"`
		TimePause     int `yaml:"timePause"`
	}
}

func getConf(file []byte) (*Config, error) {
	conf := &Config{}
	err := yaml.Unmarshal(file, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func openConfigFile() []byte {
	homePath := os.Getenv("HOME")
	file, err := ioutil.ReadFile(homePath + CONFIG_PATH)
	if err != nil {
		panic(err)
	}

	return file
}

func GetConf() (*Config, error) {
	conf, err := getConf(openConfigFile())
	if err != nil {
		return nil, err
	}

	return conf, nil
}
