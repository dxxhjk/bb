package util

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func StrSpaceTo_(str string) string {
	str = strings.Split(str, ">")[0]
	newStr := ""
	for _, c := range str {
		if c != ' ' && c != '/' {
			newStr += string(c)
		} else if c == ' ' {
			newStr += "_"
		} else if c == '/' {
			newStr += "^"
		}
	}
	return newStr
}

func OpenFile(filePath string) (*os.File, error) {
	if CheckFileIsExist(filePath) { //如果文件存在
		return os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend) //打开文件
	} else {
		return os.Create(filePath) //创建文件
	}
}

func CheckDirIsExist(dirPath string) bool {
	s, err := os.Stat(dirPath)
	if err != nil {
		return false
	}
	return s.IsDir()
}


func CheckFileIsExist(filePath string) bool {
	// 读取文件信息，判断文件是否存在
	_, err := os.Stat(filePath)
	if err != nil{
		if os.IsExist(err){  // 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}

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
