/*
Copyright © 2022 Tianjiao Huang <tjhu@tjhu.dev>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commandsGetCmd represents the get command
var commandsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list of commands. Guild commands if guild id is specified; global commands otherwise",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	commandsCmd.AddCommand(commandsGetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
