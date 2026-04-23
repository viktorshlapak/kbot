package cmd

import (
	"log"

	"kbot/internal/bot"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kbot",
	Short: "Telegram bot written in Go",
	Run: func(cmd *cobra.Command, args []string) {
		if err := bot.Start(); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}