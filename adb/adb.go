package adb

import (
	"bb/config"
	"bb/util"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

func Init(socIp string, socPortList []string) {
	// adb connect 每个 soc
	var wg sync.WaitGroup
	for _, socPort := range socPortList {
		wg.Add(1)
		adbCmdStr := "adb connect " + socIp + ":" + socPort
		go func(adbCmdStr string) {
			defer wg.Done()
			adbCmd := exec.Command("bash", "-c", adbCmdStr)
			if err := adbCmd.Run(); err != nil {
				fmt.Println("os command failed:")
				fmt.Println(adbCmdStr)
			}
		}(adbCmdStr)
	}
	wg.Wait()
	fmt.Println("finish init adb")
}

func InitSoc(socIp string, socPortList []string) {
	// 每个 soc 创建 /data/bb_workspace 文件夹
	var wg sync.WaitGroup
	for _, socPort := range socPortList {
		wg.Add(1)
		adbCmdStr := "adb -s " + socIp + ":" + socPort + " shell \"mkdir /data/bb_workspace\""
		go func(adbCmdStr string) {
			defer wg.Done()
			adbCmd := exec.Command("bash", "-c", adbCmdStr)
			if err := adbCmd.Run(); err != nil {
				fmt.Println("os command failed:")
				fmt.Println(adbCmdStr)
				fmt.Println(err)
			}
		}(adbCmdStr)
	}
	wg.Wait()
	fmt.Println("finish init adb")
}

func Push(socIp string, socPortList []string, srcFile, des string) {
	srcFile = util.GetFullPath(srcFile)
	adbCmdStr := "push " + srcFile + " " + des
	var adbCmdStrList []string
	for range socPortList {
		adbCmdStrList = append(adbCmdStrList, adbCmdStr)
	}
	execAdbCmd(socIp, socPortList, adbCmdStrList)
}

func Pull(socIp string, socPortList []string, srcFile, des string) {
	var adbCmdStrList []string
	des = util.GetFullPath(des)
	if _, err := os.Stat(des); err != nil {
		fmt.Println("Pull failed: " + des + " is not a dic.")
		return
	}
	for _, socPort := range socPortList {
		desPath := des + "/" + socPort
		if _, err := os.Stat(desPath); err != nil {
			adbCmd := exec.Command("bash", "-c", "mkdir " + desPath)
			if err = adbCmd.Run(); err != nil {
				fmt.Println("os command failed:")
				fmt.Println("mkdir " + des + "/" + socPort)
			}
		}
		adbCmdStrList = append(adbCmdStrList, "pull " + srcFile +  " " + desPath)
	}
	execAdbCmd(socIp, socPortList, adbCmdStrList)
}

func Shell(socIp string, socPortList []string, command string) {
	adbCmdStr := "shell " + command
	var adbCmdStrList []string
	for  range socPortList {
		adbCmdStrList = append(adbCmdStrList, adbCmdStr)
	}
	execAdbCmd(socIp, socPortList, adbCmdStrList)
}

func execAdbCmd(socIp string, socPortList []string, adbCmdStrList []string) {
	var wg sync.WaitGroup
	for i, socPort := range socPortList {
		wg.Add(1)
		strToExec := "adb -s " + socIp + ":" + socPort + " " + adbCmdStrList[i]
		adbLogPath := config.GetWorkPath() + "adb_log/" + socPort
		go func(adbCmdStr, strToExec, adbLogPath string) {
			defer wg.Done()
			adbCmd := exec.Command("bash", "-c", strToExec)
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			adbCmd.Stdout = &stdout
			adbCmd.Stderr = &stderr
			stdoutPath := adbLogPath + "/stdout"
			stderrPath := adbLogPath + "/stderr"
			var stdoutf *os.File
			var stderrf *os.File
			if !util.CheckDirIsExist(adbLogPath) {
				err := os.Mkdir(adbLogPath, os.ModePerm)
				if err != nil {
					fmt.Println("Error: mkdir failed." + err.Error())
					return
				}
			}
			stdoutf, err1 := util.OpenFile(stdoutPath)
			stderrf, err2 := util.OpenFile(stderrPath)
			if err1 != nil || err2 != nil {
				fmt.Println("Error: open file failed.", err1, err2)
				return
			}
			defer stdoutf.Close()
			defer stderrf.Close()
			now := time.Now()
			timeStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d",
				now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
			if err := adbCmd.Run(); err != nil {
				fmt.Println("os command failed:")
				fmt.Println(strToExec)
			}
			stdoutf.WriteString(timeStr + " # " + adbCmdStr + "\n")
			stderrf.WriteString(timeStr + " # " + adbCmdStr + "\n")
			stdoutf.WriteString(stdout.String())
			stderrf.WriteString(stderr.String())
		}(adbCmdStrList[i], strToExec, adbLogPath)
	}
	wg.Wait()
	fmt.Println("finish exec command, the first of them is:\n" + adbCmdStrList[0])
}


