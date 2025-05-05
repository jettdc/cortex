package ui

import (
	"fmt"
	"github.com/jettdc/cortex/db"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#25A065")).
		Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
		Render
)

type item struct {
	title, desc    string
	clipboardValue *db.ClipboardValue
}

func (i item) Title() string            { return i.title }
func (i item) Description() string      { return i.desc }
func (i item) FilterValue() string      { return i.title }
func (i item) Data() *db.ClipboardValue { return i.clipboardValue }

type model struct {
	list         list.Model
	delegateKeys *delegateKeyMap
}

func newModel(clipboardValues []*db.ClipboardValue) model {
	delegateKeys := newDelegateKeyMap()

	items := make([]list.Item, len(clipboardValues))
	for i, clipboardValue := range clipboardValues {
		items[i] = item{
			title:          clipboardValue.Key,
			desc:           processString(clipboardValue.Value),
			clipboardValue: clipboardValue,
		}
	}

	delegate := newItemDelegate(delegateKeys)
	groceryList := list.New(items, delegate, 0, 0)
	groceryList.Title = "Your Clipboard 📋"
	groceryList.Styles.Title = titleStyle

	return model{
		list:         groceryList,
		delegateKeys: delegateKeys,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}

func ClipboardUi(clipboardValues []*db.ClipboardValue) {
	m := newModel(clipboardValues)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func processString(input string) string {
	// Replace all newlines with the newline glyph (↵)
	input = strings.ReplaceAll(input, "\n", "\u21B5")

	return input
}
