package adb

import (
	"fmt"
	"os/exec"
	"sync"
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
	adbCmdStr := "push " + srcFile + " " + des
	var adbCmdStrList []string
	for range socPortList {
		adbCmdStrList = append(adbCmdStrList, adbCmdStr)
	}
	execAdbCmd(socIp, socPortList, adbCmdStrList)
}

func Pull(socIp string, socPortList []string, srcFile, des string) {
	adbCmdStr := "pull " + srcFile + " " + des
	var adbCmdStrList []string
	for _, socPort := range socPortList {
		adbCmdStrList = append(adbCmdStrList, adbCmdStr + "_" + socPort)
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
		go func(strToExec string) {
			defer wg.Done()
			adbCmd := exec.Command("bash", "-c", strToExec)
			if err := adbCmd.Run(); err != nil {
				fmt.Println("os command failed:")
				fmt.Println(strToExec)
			}
		}(strToExec)
	}
	wg.Wait()
	fmt.Println("finish exec command, the first of them is:\n" + adbCmdStrList[0])
}


