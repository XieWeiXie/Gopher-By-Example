package cmd

import (
	"Gopher-By-Example/graphql-example/model"
	"Gopher-By-Example/graphql-example/pkg/database"

	"github.com/spf13/cobra"
)

var SyncCMD = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s", "S", "sync", "Sync"},
	PreRun: func(cmd *cobra.Command, args []string) {
		database.MySQLDBInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		tableSync()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		database.Engine.Close()
	},
}

func tableSync() {
	collections := []interface{}{
		new(model.Vote),
		new(model.Option),
	}
	for _, i := range collections {
		database.Engine.Sync2(i)
	}
}
