package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

func GetFullPath(path string) string {
	pwdCmd := exec.Command("bash", "-c", "pwd")
	var stdout bytes.Buffer
	pwdCmd.Stdout = &stdout
	if err := pwdCmd.Run(); err != nil {
		fmt.Println("get pwd error: " + err.Error())
	}
	localPath := path
	if path[0] != '/' {
		pwd := stdout.String()
		pwd = pwd[:len(pwd) - 1] + "/"
		localPath = pwd + localPath
	}
	return localPath
}

func GetDesignatedPortList(startSocPort, socNum string, socPortList []string) ([]string, error) {
	start, _ := strconv.Atoi(startSocPort)
	num, _ := strconv.Atoi(socNum)
	socPort0, _ := strconv.Atoi(socPortList[0])
	socPortLast, _ := strconv.Atoi(socPortList[len(socPortList) - 1])
	if socPort0 > start {
		return nil, errors.New("Error: the start_soc_port must not be less than " + strconv.Itoa(socPort0))
	}
	if start + num - 1 > socPortLast {
		return nil, errors.New("Error: the last_soc_port must not be greater than " + strconv.Itoa(socPortLast))
	}
	return socPortList[start - socPort0:start - socPort0 + num], nil
}
