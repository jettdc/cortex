package ui

import (
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jettdc/cortex/db"
	"log"
	"slices"
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				item := getActiveItem(m)

				print(item.Data().Key)

				desc := item.Data().Value
				err := clipboard.WriteAll(desc)
				if err != nil {
					log.Fatal(err)
				}

				return m.NewStatusMessage(statusMessageStyle("Copied: " + title))

			case key.Matches(msg, keys.remove):
				item := getActiveItem(m)

				db.DeleteClipboardValue(item.Data().Id)

				globalIndex := getGlobalIndexForItem(item.Data().Id, m)
				m.SetItems(slices.Delete(m.Items(), globalIndex, globalIndex+1))

				if len(m.Items()) == 0 {
					keys.remove.SetEnabled(false)
				}

				return m.NewStatusMessage(statusMessageStyle("Deleted " + title))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose, keys.remove}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
	remove key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.remove,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.remove,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		remove: key.NewBinding(
			key.WithKeys("x", "backspace"),
			key.WithHelp("x", "delete"),
		),
	}
}

func getGlobalIndexForItem(id int, m *list.Model) int {
	for i, listItem := range m.Items() {
		myItem, ok := listItem.(item) // assert concrete type
		if ok && myItem.Data().Id == id {
			return i
		}
	}

	return -1
}

func getActiveItem(m *list.Model) item {
	x := m.SelectedItem()
	myItem, ok := x.(item) // assert concrete type
	if ok {
		return myItem
	}

	panic("Item not found")
}
