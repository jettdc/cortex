package main

import (
	"github.com/jettdc/cortex/cmd"
	"github.com/jettdc/cortex/db"
	"github.com/jettdc/cortex/utils"
	"log"
)

func main() {
	err := utils.EnsureCortexDir()
	if err != nil {
		log.Fatal(err)
	}

	db.InitDb()

	cmd.Execute()
}
