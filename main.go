package main

import (
	"github.com/pablotrianda/pomodox/cmd/config"
	"github.com/pablotrianda/pomodox/cmd/pomodorox"
)

func main() {
	config, err := config.GetConf()
	if err != nil {
		panic(err)
	}
	pomodorox.StartPomodorox(config)
}
