/*
Copyright © 2022 Tianjiao Huang <tjhu@tjhu.dev>

*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// True if guild specific commands.
	guild *bool
)

// commandsCmd represents the commands command
var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "Application Commands APIs. https://discord.com/developers/docs/interactions/application-commands",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		callPersistentPreRun(cmd, args)
	},
}

// commandsGetCmd represents the get command
var commandsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list of commands. Guild commands if guild id is specified; global commands otherwise",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		callPersistentPreRun(cmd, args)
		if *guild && !config.GuildID.IsValid() {
			logrus.Fatal("Guild ID empty of invalid: ", config.GuildID)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := client.Commands(config.ApplicationID)
		cobra.CheckErr(err)
		logrus.Info(commands)
	},
}

func init() {
	commandsCmd.AddCommand(commandsGetCmd)
	rootCmd.AddCommand(commandsCmd)

	// Define flags
	guild = commandsCmd.Flags().BoolP("guild", "j", false, "Guild specific command.")
}
