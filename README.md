# POMODOROX 🍅 
A pomodoro terminal app! 
<p align="center">
   <img src="https://art.pixilart.com/78565b935090b76.png" data-canonical-src="https://art.pixilart.com/78565b935090b76.png" width="400" height="350" />
</p>

![Console](https://i.imgur.com/OYp6umy.png)


## Install
1. Run `install.sh`
2. Run `$pomodorox`

## Configuration
To configure the [pomodoro](https://es.wikipedia.org/wiki/T%C3%A9cnica_Pomodoro) times and loops you must set the values on
`~/.config/pomodorox/config.yaml` \n
Disclaimer: This configurations is the way what **work for me!**, you can modify for your need. 😉
This app take 3 params: 
 * `workLoops` -> The number of block of work 
 * `timeWorkLoops` -> Time for each work block
 * `timePause` -> Time for each pause block

## Images
### Nofifications (Linux only)
![notify](https://i.imgur.com/EWRiMLk.png)

## Run on development mode
* Clone this repo `git clone https://github.com/pablotrianda/pomodorox`
* Go to folder `cd pomodorox`
* Install the dependencies `go mod tidy`
* Run `go run *.go`

# Next steps
- [ ] Show times and loops on tmux bar.  
- [ ] Use params to set the loops.
- [ ] More configurations modes.  
