/*
Copyright © 2022 Tianjiao Huang <tjhu@tjhu.dev>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commandsCmd represents the commands command
var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "Interactive with Application Commands APIs.\nhttps://discord.com/developers/docs/interactions/application-commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("commands called")
	},
}

func init() {
	rootCmd.AddCommand(commandsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
