/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"jenkins/service"
	"log"

	"github.com/spf13/cobra"
)

var project string
var branch string
var verbose bool
var force bool

var AppName = "jenkins-cli"

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(2), //最多2个自定义参数
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
	},
}

func init() {

	buildCmd.Flags().StringVarP(&project, "project", "p", "", "input server project")
	// buildCmd.MarkFlagRequired("project")

	buildCmd.Flags().StringVarP(&branch, "branch", "b", "", "input server project")
	// buildCmd.MarkFlagRequired("branch")

	buildCmd.Flags().BoolVarP(&force, "force", "f", false, "Not check for SCM changes before starting the build")

	buildCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Prints out the console output of the build.")

	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
