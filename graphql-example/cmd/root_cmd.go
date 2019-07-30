package cmd

import (
	"Gopher-By-Example/graphql-example/pkg/database"
	"Gopher-By-Example/graphql-example/pkg/router"

	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLDBInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		router.StartWebServer()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		database.Engine.Close()
	},
}

func init() {
	RootCMD.AddCommand(SyncCMD)
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		panic("command execute fail")
	}
}
