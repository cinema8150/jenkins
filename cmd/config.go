/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var url string
var user string
var token string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config gitlab server host and token",
	Long: `Config gitlab server host and token, For example:

	autotag config --url xxxx
	autotag config --user xxx
	autotag config --token xxx`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(url) > 0 {
			viper.Set("jenkins.url", url)
		}

		if len(user) > 0 {
			viper.Set("jenkins.user", user)
		}

		if len(token) > 0 {
			viper.Set("jenkins.token", token)
		}

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {

	configCmd.Flags().StringVar(&url, "url", "", "input server host")

	configCmd.Flags().StringVar(&user, "user", "", "input your name")

	configCmd.Flags().StringVar(&token, "token", "", "input your token")

	rootCmd.AddCommand(configCmd)

}
