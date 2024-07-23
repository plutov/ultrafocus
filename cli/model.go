package cli

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/plutov/ultrafocus/hosts"
)

type sessionState uint

const (
	menuView sessionState = iota
	blacklistView
)

type model struct {
	commands              []command
	commandsListSelection int
	fatalErr              error
	domains               []string
	state                 sessionState
	status                hosts.FocusStatus
	textarea              textarea.Model
}

func NewModel() model {
	domains, status, err := hosts.ExtractDomainsFromHostsFile()

	state := menuView
	ti := textarea.New()
	ti.Blur()
	if len(domains) == 0 {
		domains = hosts.DefaultDomains
	}

	return model{
		textarea: ti,
		domains:  domains,
		state:    state,
		status:   status,
		fatalErr: err,
	}
}

func (m model) Init() tea.Cmd {
	if m.fatalErr != nil {
		return tea.Quit
	}

	return nil
}

func (m *model) getCommmandsList() []command {
	if m.status == hosts.FocusStatusOn {
		return []command{commandFocusOff, commandConfigureBlacklist}
	}

	return []command{commandFocusOn, commandConfigureBlacklist}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		commands := m.getCommmandsList()
		switch msg.String() {

		case "up", "k":
			if m.state == menuView && m.commandsListSelection > 0 {
				m.commandsListSelection--
			}

		case "down", "j":
			if m.state == menuView && m.commandsListSelection < len(commands)-1 {
				m.commandsListSelection++
			}

		case "enter", " ":
			if m.state == menuView {
				m = commands[m.commandsListSelection].Run(m)
				if m.fatalErr != nil {
					return m, tea.Quit
				}
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.state == blacklistView {
				domains := strings.Split(m.textarea.Value(), "\n")
				domains = hosts.CleanDomainsList(domains)

				if err := hosts.WriteDomainsToHostsFile(domains, m.status); err != nil {
					m.fatalErr = err
					return m, tea.Quit
				}

				m.domains = domains
				m.state = menuView
				m.textarea.Blur()
			}
		}
	}

	return m, tea.Batch(cmds...)
}
