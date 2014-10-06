package main

import (
	"log"

	"github.com/kyleconroy/battleofbits/server"
	"github.com/spf13/cobra"
)

func main() {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database schema",
		Run: func(cmd *cobra.Command, args []string) {
			err := server.Migrate()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	var battleCmd = &cobra.Command{
		Use:   "battle [game] [url] [url]",
		Short: "Play a match between two players",
		Run: func(cmd *cobra.Command, args []string) {
			err := server.Battle()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "bob"}
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(battleCmd)
	rootCmd.Execute()
}
