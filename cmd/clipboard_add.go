package cmd

import (
	"fmt"
	"github.com/jettdc/cortex/db"
	"github.com/jettdc/cortex/utils/values"
	"github.com/spf13/cobra"
)

var key string
var value string

var clipboardAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new item to the clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		if len(key) == 0 {
			fmt.Println("Please provide a key")
			return
		}

		if len(value) == 0 {
			vimText, err := values.GetValueWriter().WriteValue()
			if err != nil {
				panic(err)
			}

			value = vimText
		}

		if len(value) == 0 {
			fmt.Println("Please provide a value")
			return
		}

		db.InsertClipboardValue(key, value)

		fmt.Printf("Added a new item to the clipboard: %s\n", key)
	},
}

func init() {
	clipboardAddCmd.Flags().StringVarP(&key, "key", "k", "", "The key to add")
	clipboardAddCmd.MarkFlagRequired("key")

	clipboardAddCmd.Flags().StringVarP(&value, "value", "v", "", "The value to add")

	clipboardCommand.AddCommand(clipboardAddCmd)
}
