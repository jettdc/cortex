package main

import (
	"github.com/jettdc/cortex/v2/cmd"
	"github.com/jettdc/cortex/v2/db"
	"github.com/jettdc/cortex/v2/ui"
	"github.com/jettdc/cortex/v2/utils"
	"log"
)

func main() {
	err := utils.EnsureCortexDir()
	if err != nil {
		log.Fatal(err)
	}

	db.InitDb()
	ui.InitUI()

	cmd.Execute()
}
