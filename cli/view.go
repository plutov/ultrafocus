package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var appNameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF5F87")).
	Bold(true)

var statusBarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
	Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

var statusStyle = lipgloss.NewStyle().
	Inherit(statusBarStyle).
	Foreground(lipgloss.Color("#FFFDF5")).
	Background(lipgloss.Color("#FF5F87")).
	Padding(0, 1)

var statusText = lipgloss.NewStyle().Inherit(statusBarStyle).Padding(0, 1)

var listEnumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
var listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).MarginRight(1)

var faint = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Faint(true)

func (m model) View() string {
	s := appNameStyle.Render("ultrafocus") + faint.Render(" - Reclaim your time.") + "\n\n"

	if !m.initialised {
		s += "Loading current configuration...\n\n"
	}

	if m.fatalErr != nil {
		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			statusStyle.Render("ERROR"),
			statusText.Render(m.fatalErr.Error()),
		)

		s += bar + "\n\n"
	}

	if m.initialised && m.fatalErr == nil {
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
	}

	s += "press q to quit.\n"

	return s
}
