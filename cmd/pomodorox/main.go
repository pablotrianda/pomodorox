package pomodorox

import (
	"time"
)

type Pomodorox struct {
	workLoops  timeLoop
	pauseLoops timeLoop
	workTime   time.Duration
	shortRest  time.Duration
}

type timeLoop struct {
	loops          int
	completedLoops int
}

func (p *Pomodorox) run() {
	for i := p.workLoops.loops; i != 0; i-- {
		printInfo(p)
		work(p)
		pause(p)
	}
}

func (p *Pomodorox) setCompletedWorkLoop() {
	p.workLoops.loops -= 1
	p.workLoops.completedLoops += 1
}

func (p *Pomodorox) setCompletedPauseLoop() {
	p.pauseLoops.loops -= 1
	p.pauseLoops.completedLoops += 1
}

func (p *Pomodorox) isCompleted() bool {
	return p.workLoops.loops == 0
}
