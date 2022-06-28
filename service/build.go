package service

import (
	"fmt"
	"jenkins/shell"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func Build(project string, branch string, custom bool, force bool, verbose bool) error {

	jobs, err := GetJob(project)
	if err != nil {
		return err
	}

	var job string
	if len(jobs) == 1 {
		job = jobs[0]
	} else {

		prompt := promptui.Select{
			Label: "Select Job",
			Items: jobs,
		}

		//TODO: 支持忽略大小写匹配？
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		job = result
	}

	// TODO:
	_, err = GetJobConfig(job)

	var jarFile = viper.GetString("jenkins.jar")

	var cmdStr = "java -jar " + jarFile + " -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.user") + ":" + viper.GetString("jenkins.token") + " build " + job

	var tipCmd = "jenkins build " + job

	if len(branch) > 0 {
		cmdStr += " -p branch=origin/" + branch
		tipCmd += " -b " + branch
	}

	if !force {
		cmdStr = cmdStr + " -c"
	} else {
		tipCmd += " -f"
	}

	if verbose {
		cmdStr += " -f -v"
		tipCmd += " -v"
	}

	fmt.Printf("\n✅ %s\n", tipCmd)

	res, err := shell.Exec(cmdStr)
	if err != nil {
		// log.Fatalf("jenkins build %s err: %s", job, err)
		return err
	}

	jenkinsUrl := viper.GetString("jenkins.url") + "/view/iOS/job/" + job

	res = strings.Trim(res, "\n")
	if strings.Contains(res, " - already built by ") {
		lines := strings.Split(res, "\n")
		tip := lines[len(lines)-1]
		tip = strings.Split(tip, "] ")[1]

		buildNum := strings.Split(tip, " - already built by ")[1]

		jenkinsUrl += "/" + buildNum + "/console"

		fmt.Printf("\n✅ %s\n   You can see it on: %s\n   Or force rebuild it by: %s -f\n", tip, jenkinsUrl, tipCmd)

		return nil
	}

	if len(res) == 0 {
		fmt.Printf("\n✅ You can see more build info on: %s\n", jenkinsUrl)
	}

	//TODO: format res
	fmt.Println(res)
	return nil
}
