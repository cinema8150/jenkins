package service

import (
	"errors"
	"fmt"
	"jenkins/shell"
	"strings"
	"time"

	"github.com/spf13/viper"
)

//TODO: cache与config隔离

func GetJob(project string) ([]string, error) {

	// 1. 先取缓存
	jobs, err := getLocalJob(project)
	if err != nil {
		return nil, err
	}

	if jobs == nil || len(jobs) == 0 { // 2. 缓存为空时更新
		jobs, err = getRemoteJob(project)
		if err != nil {
			return nil, err
		}
	} else {
		// 3. 定时更新
		last_fetch := viper.GetInt64("jenkins.jobs.fetch_timestamp") //FIMXE: 定义常量 //TODO: bug fetch time failure
		if last_fetch == 0 || (time.Now().Unix()-last_fetch > viper.GetInt64("jenkins.jobs.fetch_interval")) {
			_, err := getRemoteJob(project)
			if err != nil {
				return nil, err
			}
		}
	}

	if jobs == nil || len(jobs) == 0 {
		err = errors.New(fmt.Sprintf("can not found the job : %s", project))
	}

	return jobs, nil
}

func matchJob(project string, jobs []string) []string {
	var list []string
	for _, v := range jobs {
		if v == project {
			list = []string{v}
			break
		}
		if strings.Contains(v, project) {
			list = append(list, v)
		}
	}
	return list
}

func getLocalJob(project string) ([]string, error) {
	jobs := viper.GetStringSlice("jenkins.jobs.list")
	return matchJob(project, jobs), nil
}

func getRemoteJob(project string) ([]string, error) {
	jarFile := viper.GetString("jenkins.jar")
	caches, err := cacheUpdate(jarFile)
	if err != nil {
		return nil, err
	}
	return matchJob(project, caches), nil
}

func cacheUpdate(jarFile string) ([]string, error) {
	//TODO: 完善交互，如：loading……
	fmt.Println("\n✅ Update jenkins job cache")
	// 3. 本地未命中更新
	cmdStr := "java -jar " + jarFile + " -s " + viper.GetString("jenkins.url") + " -webSocket -auth " + viper.GetString("jenkins.user") + ":" + viper.GetString("jenkins.token") + " list-jobs"
	res, err := shell.Exec(cmdStr)

	if err != nil {
		return nil, err
	}

	viper.Set("jenkins.jobs.fetch_timestamp", time.Now().Unix())

	jobs := strings.Split(strings.Trim(res, "\n"), "\n")
	if len(jobs) > 0 {

		viper.Set("jenkins.jobs.list", jobs)

		err := viper.WriteConfig()
		if err != nil {
			return nil, err
		}
	}
	return jobs, nil
}

type JobParamer struct {
	Name    string
	Type    string
	Default string
}

type JobConfig struct {
	Name string
}

func GetJobConfig(project string) (*JobConfig, error) {

	return nil, errors.New(fmt.Sprintf("Unfound %s build config", project))
}
