/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"jenkins/service"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MaximumNArgs(2), //最多1个自定义参数
	Run: func(cmd *cobra.Command, args []string) {

		if err := service.CheckJenkinsConfig(); err != nil {
			log.Fatalln(err)
		}

		if len(args) > 0 {
			project = args[0]
		}
		if len(args) > 1 {
			branch = args[1]
		}

		if err := service.Build(project, branch, force, verbose); err != nil {
			log.Fatalln(err)
		}

		//TODO: usage tip
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	_, err := service.LoadConfig(AppName)
	if err != nil {
		//FIXME: log vs fmt
		log.Fatalf("load config fatal: %s", err)
		return
	}
}
