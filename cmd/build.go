/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"jenkins/shell"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var project string
var branch string
var verbose bool
var force bool

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(viper.GetString("jenkins.url")) > 0 &&
			len(viper.GetString("jenkins.user")) > 0 &&
			len(viper.GetString("jenkins.token")) > 0 {
			fmt.Println("build called")

			var cmdStr string
			if verbose {
				cmdStr = "java -jar jenkins-cli.jar -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.user") + ":" + viper.GetString("jenkins.token") + " build " + project + " -p branch=origin/" + branch + " -f -v"
			} else {
				cmdStr = "java -jar jenkins-cli.jar -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.user") + ":" + viper.GetString("jenkins.token") + " build " + project + " -p branch=origin/" + branch
			}

			if !force {
				cmdStr = cmdStr + " -c"
			}

			fmt.Println(cmdStr)
			res, err := shell.Exec(cmdStr)
			if err != nil {
				log.Fatalf("jenkins build %s err: %s", project, err)
				return
			}
			fmt.Println(res)
			return
		}

		//FIXME: log vs fmt
		log.Fatalln("need config jenkins")
	},
}

func init() {

	buildCmd.Flags().StringVarP(&project, "project", "p", "", "input server project")
	buildCmd.MarkFlagRequired("project")

	buildCmd.Flags().StringVarP(&branch, "branch", "b", "", "input server project")
	buildCmd.MarkFlagRequired("branch")

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
