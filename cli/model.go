package cli

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/plutov/ultrafocus/hosts"
)

type TickMsg time.Time

type sessionState uint

func doTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

const (
	menuView sessionState = iota
	blacklistView
	timerView
)

type model struct {
	textarea              textarea.Model
	textinput             textinput.Model
	fatalErr              error
	status                hosts.FocusStatus
	domains               []string
	commandsListSelection int
	minutesLeft           int
	state                 sessionState
}

func NewModel() model {
	domains, status, err := hosts.ExtractDomainsFromHostsFile()

	if len(domains) == 0 {
		domains = hosts.DefaultDomains
	}

	return model{
		textarea:  GetTextareaModel(),
		textinput: GetInputModel(),
		domains:   domains,
		state:     menuView,
		status:    status,
		fatalErr:  err,
	}
}

func (m model) Init() tea.Cmd {
	if m.fatalErr != nil {
		return tea.Quit
	}

	return doTick()
}

func (m *model) getCommandsList() []command {
	if m.status == hosts.FocusStatusOn {
		return []command{commandFocusOff, commandConfigureBlacklist}
	}

	return []command{commandFocusOn, commandFocusOnWithTimer, commandConfigureBlacklist}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {

	case TickMsg:
		if m.status == hosts.FocusStatusOn && m.minutesLeft > 0 {
			m.minutesLeft--
			if m.minutesLeft == 0 {
				m = commandFocusOff.Run(m)
			}
		}

		return m, doTick()
	case tea.KeyMsg:
		commands := m.getCommandsList()
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

				m.commandsListSelection = 0
				m.domains = domains
				m.state = menuView
				m.textarea.Blur()
			}
			if m.state == timerView {
				minutesStr := m.textinput.Value()
				minutes, err := strconv.Atoi(minutesStr)
				if err != nil || minutes <= 0 {
					m.fatalErr = errors.New("Invalid number of minutes")
					return m, tea.Quit
				}

				m = commandFocusOn.Run(m)

				m.minutesLeft = minutes
				m.commandsListSelection = 0
				m.state = menuView
				m.textinput.Blur()
			}
		}
	}

	return m, tea.Batch(cmds...)
}
