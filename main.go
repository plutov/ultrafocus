package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/plutov/ultrafocus/cli"
	"github.com/plutov/ultrafocus/server"
)

func main() {
	// can be started with server mode which launches the server
	if len(os.Args) > 1 && os.Args[1] == "ultrafocusserver" {
		// blocking
		server.Start()
		return
	}

	m := cli.NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("unable to launch ultrafocus: %v", err)
		os.Exit(1)
	}
}
