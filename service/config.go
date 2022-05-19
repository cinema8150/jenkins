package service

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadConfig() (string, error) {

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
