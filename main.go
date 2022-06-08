package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pablotrianda/pomodox/cmd/config"
	"github.com/pablotrianda/pomodox/cmd/pomodorox"
)

func main() {
	if !(len(os.Args) > 1) {
		fmt.Println("YOU MUST USE PARAMS!!")
		os.Exit(1)
	}

	timeWork, _ := strconv.Atoi(os.Args[1])
	timePause, _ := strconv.Atoi(os.Args[2])
	qtyWorkLoops, _ := strconv.Atoi(os.Args[3])

	conf := &config.Config{
		TimeWork:     timeWork,
		TimePause:    timePause,
		QtyWorkLoops: qtyWorkLoops,
	}

	pomodorox.StartPomodorox(conf)
}
