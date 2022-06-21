package service

import (
	"errors"

	"github.com/spf13/viper"
)

func CheckJenkinsConfig() error {
	if len(viper.GetString("jenkins.url")) > 0 &&
		len(viper.GetString("jenkins.user")) > 0 &&
		len(viper.GetString("jenkins.token")) > 0 {
		return nil
	}
	return errors.New("need config jenkins")
}
