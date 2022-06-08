package pomodorox

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/pablotrianda/pomodox/cmd/config"
)

func work(p *Pomodorox) {
	notifyToDesktop("Lets work!")
	time.Sleep(p.workTime)
	p.setCompletedWorkLoop()
}

func pause(p *Pomodorox) {
	notifyToDesktop("Take a break!")
	time.Sleep(p.shortRest)
	p.setCompletedPauseLoop()
	if p.isCompleted() {
		notifyToDesktop("Finished")
	}
}

func notifyToDesktop(msg string) {
	cmd := exec.Command("notify-send", "-i", "tomato", "POMODOROX", msg, "-u", "CRITICAL")
	cmd.Run()
}

func printInfo(pomo *Pomodorox) {
	fmt.Print("\033[H\033[2J") // Clean screen
	fmt.Println("🍅 POMODOROX 🍅")
	fmt.Printf("WORK TIME: %v\n", pomo.workTime)
	fmt.Printf("PAUSE TIME: %v\n", pomo.shortRest)
	totalLoops := pomo.workLoops.completedLoops + pomo.workLoops.loops
	fmt.Printf("WORKED LOOPS: %v/%v\n", pomo.workLoops.completedLoops, totalLoops)
}

func StartPomodorox(config *config.Config) {
	work := timeLoop{
		loops:          config.QtyWorkLoops,
		completedLoops: 0,
	}

	pause := timeLoop{
		loops:          config.QtyWorkLoops - 1,
		completedLoops: 0,
	}

	workTime := time.Duration(config.TimeWork)
	pauseTime := time.Duration(config.TimePause)

	pomo := Pomodorox{
		workLoops:  work,
		pauseLoops: pause,
		workTime:   time.Minute * workTime,
		shortRest:  time.Minute * pauseTime,
	}

	pomo.run()
}
