package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/plutov/ultrafocus/cli"
)

func main() {
	m := cli.NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("unable to launch ultrafocus: %v", err)
		os.Exit(1)
	}
}
