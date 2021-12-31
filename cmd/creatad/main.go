package main

import (
	"os"

	"github.com/creatachain/creata-sdk/server"
	svrcmd "github.com/creatachain/creata-sdk/server/cmd"

	app "github.com/creatachain/gaia/v4/app"
	"github.com/creatachain/gaia/v4/cmd/creatad/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
