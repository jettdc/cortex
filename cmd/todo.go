package cmd

import (
	"fmt"
	"github.com/jettdc/cortex/v2/db"
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
		for _, result := range results {
			fmt.Println(result)
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
