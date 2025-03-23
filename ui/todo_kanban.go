package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/jettdc/cortex/v2/db"
	"github.com/rivo/tview"
)

var cursorX int = 0
var itemIsActivated bool = false

type TodoKanban struct {
	items  []db.Todo
	onSave func([]db.Todo)
}

func NewTodoKanban(items []db.Todo, onSave func([]db.Todo)) *TodoKanban {
	return &TodoKanban{
		items:  items,
		onSave: onSave,
	}
}

func (kanban *TodoKanban) Render() {
	todoList := tview.NewList().ShowSecondaryText(false)
	todoList.SetBorder(true).SetTitle("TODO")
	for i, result := range filterItemsByStatus(kanban.items, "todo") {
		todoList.AddItem(getDefaultTodoItemText(result), "", rune(i+1), nil)
	}

	doingList := tview.NewList().ShowSecondaryText(false)
	doingList.SetBorder(true).SetTitle("Doing")
	for i, result := range filterItemsByStatus(kanban.items, "doing") {
		doingList.AddItem(getDefaultTodoItemText(result), "", rune(i+1), nil)
	}
	doingList.SetSelectedStyle(tcell.Style{}.Background(tcell.ColorBlack))
	doingList.SetSelectedStyle(tcell.Style{}.Foreground(tcell.ColorWhite))

	doneList := tview.NewList().ShowSecondaryText(false)
	doneList.SetBorder(true).SetTitle("Done")
	for i, result := range filterItemsByStatus(kanban.items, "done") {
		doneList.AddItem(getDefaultTodoItemText(result), "", rune(i+1), nil)
	}
	doneList.SetSelectedStyle(tcell.Style{}.Background(tcell.ColorBlack))
	doneList.SetSelectedStyle(tcell.Style{}.Foreground(tcell.ColorWhite))

	flex := tview.NewFlex().
		AddItem(todoList, 0, 1, true).
		AddItem(doingList, 0, 1, true).
		AddItem(doneList, 0, 1, false)

	withFooter := tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 0, 0)
	setDefaultFooter(withFooter)

	GetApp().SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Rune() == 'q' {
			GetApp().Stop()
			return nil
		}
		if itemIsActivated {
			if event.Key() == tcell.KeyEnter {
				deactivateItemFromList(todoList, doingList, doneList)
				setDefaultFooter(withFooter)
			} else if event.Key() == tcell.KeyUp {
				moveSelectedUp(todoList, doingList, doneList)
			} else if event.Key() == tcell.KeyDown {
				moveSelectedDown(todoList, doingList, doneList)
			} else if event.Key() == tcell.KeyRight {
				moveSelectedRight(todoList, doingList, doneList)
			} else if event.Key() == tcell.KeyLeft {
				moveSelectedLeft(todoList, doingList, doneList)
			}

			return nil
		} else {
			if event.Key() == tcell.KeyRight {
				incrementX(todoList, doingList, doneList)
			} else if event.Key() == tcell.KeyLeft {
				decrementX(todoList, doingList, doneList)
			} else if event.Key() == tcell.KeyEnter {
				activateItemFromList(todoList, doingList, doneList)
				setActivatedFooter(withFooter)
			}
		}

		return event
	})

	if err := GetApp().SetRoot(withFooter, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func filterItemsByStatus(items []db.Todo, status string) []db.Todo {
	newItems := make([]db.Todo, 0)
	for _, item := range items {
		if item.Status == status {
			newItems = append(newItems, item)
		}
	}

	return newItems
}

func getDefaultTodoItemText(item db.Todo) string {
	formatted := fmt.Sprintf("[P%d] %s", item.Priority, item.Message)
	return tview.Escape(formatted)
}

func incrementX(lists ...*tview.List) {
	if len(lists) == 0 {
		return
	}

	if cursorX >= len(lists)-1 {
		return
	}

	prevList := lists[cursorX]

	cursorX += 1

	for i, curList := range lists {
		if i == cursorX {
			setSelectedStyleDefault(curList)
			GetApp().SetFocus(curList)
			curList.SetCurrentItem(prevList.GetCurrentItem())
		} else {
			setSelectedStyleNone(curList)
		}
	}
}

func decrementX(lists ...*tview.List) {
	if len(lists) == 0 {
		return
	}

	if cursorX <= 0 {
		return
	}

	prevList := lists[cursorX]
	cursorX -= 1

	for i, curList := range lists {
		if i == cursorX {
			setSelectedStyleDefault(curList)
			GetApp().SetFocus(curList)
			curList.SetCurrentItem(prevList.GetCurrentItem())
		} else {
			setSelectedStyleNone(curList)
		}
	}
}

func setSelectedStyleNone(list *tview.List) {
	list.SetSelectedBackgroundColor(tcell.ColorBlack)
	list.SetSelectedTextColor(tcell.ColorWhite)
}

func setSelectedStyleDefault(list *tview.List) {
	list.SetSelectedBackgroundColor(tcell.ColorWhite)
	list.SetSelectedTextColor(tcell.ColorBlack)
}

func activateItemFromList(lists ...*tview.List) {
	listToActivateIn := lists[cursorX]

	selectedIndex := listToActivateIn.GetCurrentItem()
	toActivateText, _ := listToActivateIn.GetItemText(selectedIndex)
	listToActivateIn.SetItemText(selectedIndex, getActivatedText(toActivateText), "")

	itemIsActivated = true
}

func getActivatedText(unactivatedText string) string {
	return fmt.Sprintf("[white:blue]%s[white:blue]", unactivatedText)
}

func deactivateItemFromList(lists ...*tview.List) {
	listToActivateIn := lists[cursorX]

	selectedIndex := listToActivateIn.GetCurrentItem()
	toDeactivateText, _ := listToActivateIn.GetItemText(selectedIndex)
	listToActivateIn.SetItemText(selectedIndex, removeColorTags(toDeactivateText), "")

	itemIsActivated = false
}

func removeColorTags(text string) string {
	// Ensure the string is long enough to slice
	if len(text) > 2 {
		// Remove the first and last 14 characters which are "[white:blue]"
		// (adjust this based on your tag length if needed)
		return text[12 : len(text)-12]
	}
	return text
}

// TODO(jcrowson): Only move up/down within priority

func moveSelectedUp(lists ...*tview.List) {
	listToActivateIn := lists[cursorX]

	selectedIndex := listToActivateIn.GetCurrentItem()
	selectedText, _ := listToActivateIn.GetItemText(selectedIndex)
	listToActivateIn.RemoveItem(selectedIndex)
	listToActivateIn.InsertItem(selectedIndex-1, selectedText, "", rune(1), nil)
	listToActivateIn.SetCurrentItem(selectedIndex - 1)
}

func moveSelectedDown(lists ...*tview.List) {
	listToActivateIn := lists[cursorX]

	selectedIndex := listToActivateIn.GetCurrentItem()
	selectedText, _ := listToActivateIn.GetItemText(selectedIndex)

	// Check if the selected index is not the last item
	if selectedIndex < listToActivateIn.GetItemCount()-1 {
		listToActivateIn.RemoveItem(selectedIndex)
		listToActivateIn.InsertItem(selectedIndex+1, selectedText, "", rune(1), nil)
		listToActivateIn.SetCurrentItem(selectedIndex + 1)
	}
}

func moveSelectedRight(lists ...*tview.List) {
	if cursorX >= len(lists)-1 {
		return
	}

	moveFrom := lists[cursorX]
	moveFromSelectedIndex := moveFrom.GetCurrentItem()
	moveFromSelectedText, _ := moveFrom.GetItemText(moveFromSelectedIndex)
	moveFrom.RemoveItem(moveFromSelectedIndex)

	moveTo := lists[cursorX+1]
	moveTo.InsertItem(moveFromSelectedIndex, moveFromSelectedText, "", rune(1), nil)
	incrementX(lists...)

	moveTo.SetCurrentItem(moveFromSelectedIndex)

}

func moveSelectedLeft(lists ...*tview.List) {
	if cursorX <= 0 {
		return
	}

	// Get the list from which the item is being moved
	moveFrom := lists[cursorX]
	moveFromSelectedIndex := moveFrom.GetCurrentItem()
	moveFromSelectedText, _ := moveFrom.GetItemText(moveFromSelectedIndex)

	// Remove the selected item from the current list
	moveFrom.RemoveItem(moveFromSelectedIndex)

	// Get the list to the left (cursorX-1)
	moveTo := lists[cursorX-1]
	moveTo.InsertItem(moveFromSelectedIndex, moveFromSelectedText, "", rune(1), nil)

	// Decrease cursorX to move left
	decrementX(lists...)

	// Set the current item in the list to the one that was moved
	moveTo.SetCurrentItem(moveFromSelectedIndex)
}

func setDefaultFooter(frame *tview.Frame) {
	frame.Clear()
	frame.AddText("[↑ - Up] [↓ - Down] [← - Left] [→ - Right] [ENTER - Select] [q - Quit]", false, tview.AlignBottom, tcell.ColorWhite)
}

func setActivatedFooter(frame *tview.Frame) {
	frame.Clear()
	frame.AddText("[↑ - Move up] [↓ - Move down] [← - Move left] [→ - Move right] [ENTER - Deselect] [e - Edit] [p - Set priority] [q - Quit]", false, tview.AlignBottom, tcell.ColorWhite)
}

//
//import "github.com/rivo/tview"
//
//type KanbanColumnConfig struct {
//}
//
//type KanbanConfig[T any] struct {
//	columns []KanbanColumnConfig
//	OnSave  func(KanbanColumnResult) error
//}
//
//func NewKanban[T any](config KanbanConfig[T]) *tview.Flex {
//	xxx := 2
//}
//main, _ := todoList.GetItemText(0)
//	todoList.SetItemText(0, fmt.Sprintf("[white:blue]%s[white:blue]", main), "")
