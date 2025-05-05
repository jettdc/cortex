package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jettdc/cortex/v2/db"
	"log"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)
)

type model struct {
	columns    []kanbanColumn
	cursorX    int
	cursorY    int
	isSelected bool
}

type kanbanColumn struct {
	name  string
	list list.Model
}

type todoItem struct {
	data db.Todo
}

func (t todoItem) FilterValue() string {
	return t.data.Message
}

func initModel(todos []db.Todo) model {
	columnNames := getUniqueStatuses(todos)

	columns := make([]kanbanColumn, len(columnNames))
	for i, columnName := range columnNames {
		filteredItems := filterItemsByStatus(columnName, todos)
		listItems := make([]todoItem, len(filteredItems))
		for _, item := range filteredItems {
			listItems = append(listItems, todoItem{data: item})
		}

		columnList := list.New(listItems)

		columns[i] = kanbanColumn{
			name:  columnName,
			list: columnList,
		}
	}

	i := make([]list.Item, 0)
	i[0] = list.Item()
	return model{
		columns: []kanbanColumn{
			kanbanColumn{
				name:  "TODO",
				items: ,
			},
			kanbanColumn{
				name:  "Doing",
				items: make([]db.Todo, 0),
			},
			kanbanColumn{
				name:  "Done",
				items: make([]db.Todo, 0),
			},
		},
		cursorX: 0,
		cursorY: 0,
	}
}

func getUniqueStatuses(todos []db.Todo) []string {
	statuses := make(map[string]bool)
	for _, todo := range todos {
		statuses[todo.Status] = true
	}

	statusNames := make([]string, len(statuses))
	for key, _ := range statuses {
		statusNames = append(statusNames, key)
	}

	return statusNames
}

func filterItemsByStatus(status string, todos []db.Todo) []db.Todo {
	filtered := make([]db.Todo, 0)
	for _, todo := range todos {
		if todo.Status != status {
			continue
		}

		filtered = append(filtered, todo)
	}
	return filtered
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursorY > 0 {
				m.cursorY--
			}

		case "down", "j":
			if m.cursorY < len(m.columns[m.cursorX].items)-1 {
				m.cursorY++
			}

		case "enter", " ":
			m.isSelected = true
		}
	}

	return m, nil
}

func (m model) View() string {
	return "testing1"
}

func Kanban() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
