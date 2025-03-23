package cmd

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/jettdc/cortex/v2/db"
	"github.com/jettdc/cortex/v2/ui"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"os"
)

var message string
var priority int8

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Prints a greeting message",
	Run: func(cmd *cobra.Command, args []string) {
		messageError := validateMessage(message)
		if messageError != nil {
			fmt.Println(messageError.Error())
			os.Exit(1)
		}

		priorityError := validatePriority(priority)
		if priorityError != nil {
			fmt.Println(priorityError.Error())
			os.Exit(1)
		}

		db.InsertTodo(message, priority)
	},
}

var todoLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Gets all the todos",
	Run: func(cmd *cobra.Command, args []string) {
		results := db.GetTodos()

		list := tview.NewList().ShowSecondaryText(false)
		list.SetBorder(true).SetTitle("TODO")

		list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyRune || event.Rune() == 'q' {
				// Exit the application on Escape
				ui.GetApp().Stop()
			}
			return event
		})

		list2 := tview.NewList().ShowSecondaryText(false)
		list2.SetBorder(true).SetTitle("Doing")

		list3 := tview.NewList().ShowSecondaryText(false)
		list3.SetBorder(true).SetTitle("Done")

		for i, result := range results {
			tx := fmt.Sprintf("[P%d] %s", result.Priority, result.Message)
			list.AddItem(tview.Escape(tx), "", rune(i+1), nil)
		}

		list.SetCurrentItem(0)
		main, _ := list.GetItemText(0)
		list.SetItemText(0, fmt.Sprintf("[white:green]%s[white:green]", main), "")

		flex := tview.NewFlex().
			AddItem(list, 0, 1, true).
			AddItem(list2, 0, 1, true).
			AddItem(list3, 0, 1, false)

		if err := ui.GetApp().SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	// root
	rootCmd.AddCommand(todoCmd)
	todoCmd.Flags().StringVarP(&message, "message", "m", "", "TODO message")
	todoCmd.Flags().Int8VarP(&priority, "priority", "p", 3, "Priority")

	// todo __
	todoCmd.AddCommand(todoLsCmd)
}

func validateMessage(message string) error {
	if message == "" {
		return fmt.Errorf("TODO message must not be empty")
	}

	return nil
}

func validatePriority(priority int8) error {
	if priority < 0 || priority > 3 {
		return fmt.Errorf("priority must be between 0 and 3")
	}

	return nil
}
