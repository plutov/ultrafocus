package cli

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/plutov/ultrafocus/hosts"
)

type model struct {
	hostsManager          *hosts.Manager
	commands              []command
	commandsListSelection int
	currentCommand        *command
	initialised           bool
	fatalErr              error
	domains               []string
	isEditingDomains      bool
	status                hosts.FocusStatus
	textarea              textarea.Model
}

func NewModel() model {
	return model{
		hostsManager: &hosts.Manager{},
		commands:     []command{},
	}
}

func (m model) Init() tea.Cmd {
	return m.loadInitialConfig
}

type initResult struct {
	err error
}

func (m model) loadInitialConfig() tea.Msg {
	initErr := m.hostsManager.Init()

	return initResult{
		err: initErr,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case initResult:
		m.initialised = true

		if msg.err != nil {
			m.fatalErr = msg.err
			return m, tea.Quit
		}

		m.domains = m.hostsManager.Domains
		m.status = m.hostsManager.Status
		if m.status == hosts.FocusStatusOn {
			m.commands = []command{commandFocusOff, commandConfigureBlacklist}
		} else {
			m.commands = []command{commandFocusOn, commandConfigureBlacklist}
		}
		if len(m.domains) == 0 {
			ti := textarea.New()
			ti.Placeholder = "Once upon a time..."
			ti.Focus()
			m.textarea = ti
			m.isEditingDomains = true
			return m, textarea.Blink
		}

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if m.commandsListSelection > 0 {
				m.commandsListSelection--
			}

		case "down", "j":
			if m.commandsListSelection < len(m.commands)-1 {
				m.commandsListSelection++
			}

		case "enter", " ":
			m.currentCommand = &m.commands[m.commandsListSelection]
			return m, m.currentCommand.Run()

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
