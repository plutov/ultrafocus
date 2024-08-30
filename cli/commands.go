package cli

import (
	"strings"

	"github.com/plutov/ultrafocus/hosts"
	"github.com/plutov/ultrafocus/server"
)

type command struct {
	Run  func(m model) model
	Name string
	Desc string
}

var commandFocusOn = command{
	Name: "focus on",
	Desc: "Start focus window.",
	Run: func(m model) model {
		if err := hosts.WriteDomainsToHostsFile(m.domains, hosts.FocusStatusOn); err != nil {
			m.fatalErr = err
			return m
		}

		go server.StartAsSubprocess()
		m.status = hosts.FocusStatusOn
		m.minutesLeft = 0
		return m
	},
}

var commandFocusOff = command{
	Name: "focus off",
	Desc: "Stop focus window.",
	Run: func(m model) model {
		if err := hosts.WriteDomainsToHostsFile(m.domains, hosts.FocusStatusOff); err != nil {
			m.fatalErr = err
			return m
		}

		go server.StopSubprocess()
		m.status = hosts.FocusStatusOff
		m.minutesLeft = 0
		return m
	},
}

var commandConfigureBlacklist = command{
	Name: "blacklist",
	Desc: "Configure blacklist.",
	Run: func(m model) model {
		m.state = blacklistView
		m.textarea.SetValue(strings.Join(m.domains, "\n"))
		m.textarea.Focus()
		m.textarea.CursorEnd()
		return m
	},
}

var commandFocusOnWithTimer = command{
	Name: "focus on (timer)",
	Desc: "Start timed focus window.",
	Run: func(m model) model {
		m.state = timerView
		m.textinput.SetValue("30")
		m.textinput.Focus()
		return m
	},
}
