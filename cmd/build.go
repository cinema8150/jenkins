/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"jenkins/shell"
	"log"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	Run: func(cmd *cobra.Command, args []string) {
		if len(viper.GetString("jenkins.url")) > 0 &&
			len(viper.GetString("jenkins.user")) > 0 &&
			len(viper.GetString("jenkins.token")) > 0 {

			var jarFile = viper.GetString("jenkins.jar")
			//TODO: bug 升级版本后地址会发生变更
			if len(jarFile) == 0 {

				info, err := shell.Exec("brew info " + AppName)
				if err != nil {
					log.Fatal(err)
				}

				infos := strings.Split(info, "\n")
				for _, v := range infos {
					if strings.Contains(v, "/"+AppName+"/") {
						jarFile = path.Join(strings.Split(v, " ")[0], "libexec/bin/jenkins-cli.jar")
						viper.Set("jenkins.jar", jarFile)
						viper.WriteConfig()
						break
					}
				}

			}

			// 1. 本地查询
			var list []string

			jobs := viper.GetStringSlice("jenkins.jobs")
			if jobs == nil {
				//TODO: 完善交互，如：loading……
				fmt.Println("query jenkons job")
				// 3. 本地未命中更新
				cmdStr := "java -jar " + jarFile + " -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.user") + ":" + viper.GetString("jenkins.token") + " list-jobs"
				res, err := shell.Exec(cmdStr)

				if err != nil {
					log.Fatalln(err)
				}

				jobs = strings.Split(strings.Trim(res, "\n"), "\n")
				if len(jobs) > 0 {

					viper.Set("jenkins.jobs", jobs)

					err := viper.WriteConfig()
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				//TODO: 定时更新
			}

			for _, v := range jobs {
				if v == project {
					list = []string{v}
					break
				}
				if strings.Contains(v, project) {
					list = append(list, v)
				}
			}

			if list == nil {
				log.Fatalf("can not found the job : %s", project)
				return
			}

			var job string
			if len(list) == 1 {
				job = list[0]
			} else {

				prompt := promptui.Select{
					Label: "Select Job",
					Items: list,
				}

				//TODO: 支持忽略大小写匹配？
				_, result, err := prompt.Run()

				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}

				job = result
			}

			var cmdStr string
			if verbose {
				cmdStr = "java -jar " + jarFile + " -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.user") + ":" + viper.GetString("jenkins.token") + " build " + job + " -p branch=origin/" + branch + " -f -v"
			} else {
				cmdStr = "java -jar " + jarFile + " -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.user") + ":" + viper.GetString("jenkins.token") + " build " + job + " -p branch=origin/" + branch
			}

			if !force {
				cmdStr = cmdStr + " -c"
			}

			res, err := shell.Exec(cmdStr)
			if err != nil {
				log.Fatalf("jenkins build %s err: %s", job, err)
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
