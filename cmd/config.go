/*
Copyright © 2022 Tianjiao Huang <tjhu@tjhu.dev>

*/
package cmd

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and persist configurations.",
	Long:  `Set and persist configurations(discord token, application, etc.) to the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.ConfigFileUsed() == "" {
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)
			default_config_path := path.Join(home, CONFIG_FILENAME)
			viper.SetConfigFile(default_config_path)
			logrus.Warn("No existing config file found. Writing configs to default path ", default_config_path)
		}
		cobra.CheckErr(viper.WriteConfig())
		logrus.Info("Configs written to file ", viper.ConfigFileUsed())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
