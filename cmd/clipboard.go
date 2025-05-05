package cmd

import (
	"github.com/jettdc/cortex/db"
	"github.com/jettdc/cortex/ui"
	"github.com/spf13/cobra"
)

var clipboardCommand = &cobra.Command{
	Use:     "clipboard",
	Aliases: []string{"cb"},
	Short:   "Quick access clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		values := db.GetAllClipboardValues()
		ui.ClipboardUi(values)
	},
}

func init() {
	rootCmd.AddCommand(clipboardCommand)
}
