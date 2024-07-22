package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var appNameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF5F87")).
	Bold(true)

var errorAlertStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFDF5")).
	Background(lipgloss.Color("#FF5F87")).
	Padding(0, 1)

var errorInfoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("250")).
	Padding(0, 1)

var listEnumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
var listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).MarginRight(1)

var faint = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Faint(true)

func (m model) View() string {
	s := appNameStyle.Render("ultrafocus") + faint.Render(" - Reclaim your time.") + "\n\n"

	if !m.initialised {
		s += faint.Render("...") + "\n"
		return s
	}

	if m.fatalErr != nil {
		s += errorAlertStyle.Render("ERROR") + errorInfoStyle.Render(m.fatalErr.Error()) + "\n"
		return s
	}

	l := list.New().Enumerator(func(items list.Items, i int) string {
		if i == m.commandsListSelection {
			return "→"
		}
		return " "
	}).
		EnumeratorStyle(listEnumeratorStyle).
		ItemStyle(listItemStyle)
	for _, c := range m.commands {
		l.Item(c.Name + faint.Render(" - "+c.Desc))
	}
	s += l.String() + "\n\n"

	s += "press q to quit.\n"

	return s
}
