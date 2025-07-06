package server

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/process"
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
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>Ultra Focus</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            background-color: #f0f0f0;
        }
        .message {
            background-color: white;
            padding: 2rem;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            text-align: center;
            font-size: 1.2rem;
            max-width: 600px;
        }
    </style>
</head>
<body>
    <div class="message">
        <h1>Ultra Focus</h1>
        <p>%s</p>
    </div>
</body>
</html>`, msg)
}

func Start() {
	http.HandleFunc("/", focusMessage)

	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return
	}
	cfg := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	srv := &http.Server{
		TLSConfig:    cfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	go srv.ListenAndServeTLS("", "")
	http.ListenAndServe("127.0.0.1:80", nil)
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
