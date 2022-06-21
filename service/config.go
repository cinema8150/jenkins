package service

import (
	"jenkins/shell"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadConfig(app string) (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	var (
		cfgFileFolder = path.Join(home, ".cinema8150/jenkins-build")
		cfgFileName   = ".jenkins"
		cfgFileType   = "yaml"
		cfgFilePath   = path.Join(cfgFileFolder, cfgFileName+"."+cfgFileType)
	)

	cobra.CheckErr(err)
	viper.SetConfigName(cfgFileName)
	viper.SetConfigType(cfgFileType)
	viper.AddConfigPath(cfgFileFolder)

	//FIXME:
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); ok { // config not found

		os.MkdirAll(cfgFileFolder, os.ModePerm)
		os.Create(cfgFilePath)

		err = setDefaultConfigs(app)
		if err != nil {
			return "", err
		}

		err = viper.WriteConfig()
		if err != nil {
			return "", err
		}
	} else if err != nil { // unkonwn err
		return "", err
	}

	viper.ConfigFileUsed()

	return cfgFilePath, nil

}

func setDefaultConfigs(app string) error {

	// 1. set jenkins.jar path
	//TODO: bug 升级版本后地址会发生变更
	info, err := shell.Exec("brew info " + app)
	if err != nil {
		return err
	}

	infos := strings.Split(info, "\n")
	for _, v := range infos {
		if strings.Contains(v, "/"+app+"/") {
			jarFile := path.Join(strings.Split(v, " ")[0], "libexec/bin/jenkins-cli.jar")
			viper.Set("jenkins.jar", jarFile)
			break
		}
	}

	// 2. set jenkins.jobs.fetch_interval(s)
	viper.Set("jenkins.jobs.fetch_interval", 60*60*24)

	// 3.

	return nil

}
