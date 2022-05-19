package shell

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

func init() {

}

// Excute 执行cmd指令并返回执行状态和结果信息
func Exec(cmd string) (string, error) {
	branches := strings.Split(cmd, " ")
	CMD := exec.Command(branches[0], branches[1:]...)
	CMD.Env = os.Environ()
	CMD.Env = append(CMD.Env, "LANG=en_US.UTF-8")
	var stdout, stderr bytes.Buffer
	CMD.Stdout = &stdout
	CMD.Stderr = &stderr
	err := CMD.Run()
	outStr, errStr := stdout.String(), stderr.String()
	if err != nil {
		return "", errors.New(errStr)
	}
	return outStr, nil
}

func ExecOnDir(cmd string, path string) (string, error) {
	branches := strings.Split(cmd, " ")
	CMD := exec.Command(branches[0], branches[1:]...)
	CMD.Env = os.Environ()
	CMD.Env = append(CMD.Env, "LANG=en_US.UTF-8")
	CMD.Dir = path
	var stdout, stderr bytes.Buffer
	CMD.Stdout = &stdout
	CMD.Stderr = &stderr
	err := CMD.Run()
	outStr, errStr := stdout.String(), stderr.String()
	if err != nil {
		return "", errors.New(errStr)
	}
	return outStr, nil
}

func ExecWithOption(cmd string) (string, error) {
	CMD := exec.Command("bash", "-c", cmd)
	CMD.Env = os.Environ()
	CMD.Env = append(CMD.Env, "LANG=en_US.UTF-8")
	var stdout, stderr bytes.Buffer
	CMD.Stdout = &stdout
	CMD.Stderr = &stderr
	err := CMD.Run()
	outStr, errStr := stdout.String(), stderr.String()
	if err != nil {
		return "", errors.New(errStr)
	}
	return outStr, nil
}

func ExecWithOptionOnDir(cmd string, path string) (string, error) {
	// branches := strings.Split(cmd, " ")
	CMD := exec.Command("bash", "-c", cmd)
	CMD.Env = os.Environ()
	CMD.Env = append(CMD.Env, "LANG=en_US.UTF-8")
	CMD.Dir = path
	var stdout, stderr bytes.Buffer
	CMD.Stdout = &stdout
	CMD.Stderr = &stderr
	err := CMD.Run()
	outStr, errStr := stdout.String(), stderr.String()
	if err != nil {
		return "", errors.New(errStr)
	}
	return outStr, nil
}
