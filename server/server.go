package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

var focusMessages = []string{
	"Mind over matter. Conquer your day.",
	"One task at a time. Progress, not perfection.",
	"You've got this. Believe in your abilities.",
	"Small steps, big results. Keep moving forward.",
	"Focus is key. Unlock your potential.",
	"Stay positive, stay focused. Good vibes only.",
	"Prioritize and conquer. Make every minute count.",
	"Clear your mind, clear your path. Find your zen.",
	"Limit distractions, maximize output. Stay on track.",
	"Your future self will thank you. Invest in today.",
}

func focusMessage(w http.ResponseWriter, r *http.Request) {
	msg := focusMessages[rand.Intn(len(focusMessages))]
	fmt.Fprintf(w, msg)
}

func Start() {
	http.HandleFunc("/", focusMessage)

	go http.ListenAndServe(":443", nil)
	http.ListenAndServe(":80", nil)
}

func StartAsSubprocess() {
	ex, err := os.Executable()
	if err != nil {
		return
	}

	cmd := exec.Command(ex, "ultrafocusserver")
	err = cmd.Start()
	if err != nil {
		return
	}
}

func StopSubprocess() {
	processes, err := process.Processes()
	if err != nil {
		return
	}
	for _, p := range processes {
		line, _ := p.Cmdline()
		if strings.Contains(line, "ultrafocusserver") {
			p.Kill()
			return
		}
	}
}
