package adb

import (
	"bb/config"
	"bb/util"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	fmt.Println("adb init finished")
}

func InitSocWorkSpace(socIp string, socPortList []string) {
	// 每个 soc 创建 /data/bb_workspace 文件夹
	Shell(socIp, socPortList, "mkdir /data/bb_workspace", false, "")
}

func Push(socIp string, socPortList []string, srcFile, des string) {
	srcFile = util.GetFullPath(srcFile)
	adbCmdStr := "push " + srcFile + " " + des
	var adbCmdStrList []string
	for range socPortList {
		adbCmdStrList = append(adbCmdStrList, adbCmdStr)
	}
	execAdbCmd(socIp, socPortList, adbCmdStrList, false)
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
	execAdbCmd(socIp, socPortList, adbCmdStrList, false)
}

func Shell(socIp string, socPortList []string, command string, energy bool, energyMonitorOutput string) {
	adbCmdStr := "shell \"" + command + "\""
	var adbCmdStrList []string
	for  range socPortList {
		adbCmdStrList = append(adbCmdStrList, adbCmdStr)
	}
	if energy {
		socPortList, adbCmdStrList = addBmc(socPortList, adbCmdStrList, command, energyMonitorOutput)
	}
	execAdbCmd(socIp, socPortList, adbCmdStrList, energy)
}

func addBmc(socPortList, adbCmdStrList []string, command string, energyMonitorOutput string) ([]string, []string) {
	if energyMonitorOutput == "" {
		command = util.StrSpaceTo_(command)
		now := time.Now()
		timeStr := fmt.Sprintf("%02d-%02d-%02d-%02d:%02d:%02d",
			now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
		socPortList = append([]string{config.GetBmcPort()}, socPortList...)
		bmcCmdStr := "/root/bmc_batch_usage_monitor.sh " + command + "_" + timeStr + " &"
		adbCmdStrList = append([]string{bmcCmdStr}, adbCmdStrList...)
	}  else {
		socPortList = append([]string{config.GetBmcPort()}, socPortList...)
		bmcCmdStr := "/root/bmc_batch_usage_monitor.sh " + energyMonitorOutput + " &"
		adbCmdStrList = append([]string{bmcCmdStr}, adbCmdStrList...)
	}
	return socPortList, adbCmdStrList
}

func execAdbCmd(socIp string, socPortList []string, adbCmdStrList []string, energy bool) {
	var wg sync.WaitGroup
	var bmcPort string
	var energyMonitorCmdStr string
	var bmcPid string
	var bmcPidCh chan string
	if energy {
		bmcPort, energyMonitorCmdStr = socPortList[0], adbCmdStrList[0]
		socPortList, adbCmdStrList = socPortList[1:], adbCmdStrList[1:]
		getPidCmdStr := "ssh -p " + bmcPort + " root@" + socIp + " \"" + "ps -ef | grep \"/root/bmc_batch_usage_monitor.sh\" | grep -v grep\""
		energyMonitorCmdStr = "ssh -p " + bmcPort + " root@" + socIp + " \"" + energyMonitorCmdStr + "\" &"
		bmcPidCh = make(chan string)
		go func(bmcCmdStr, getPidCmdStr string, ch chan string) {
			energyMonitorCmd := exec.Command("bash", "-c", bmcCmdStr)
			getPidCmd := exec.Command("bash", "-c", getPidCmdStr)
			var stdout bytes.Buffer
			getPidCmd.Stdout = &stdout
			if err := energyMonitorCmd.Run(); err != nil {
				fmt.Println("start energy monitor failed.")
			}
			time.Sleep(3000 * time.Millisecond)
			if err := getPidCmd.Run(); err != nil {
				fmt.Println("get energy monitor pid failed.")
			}
			bmcPidCh <- stdout.String()
		}(energyMonitorCmdStr, getPidCmdStr,bmcPidCh)
		bmcPid = <-bmcPidCh
		for _, bmcSplitStr := range strings.Split(bmcPid, " ") {
			if bmcSplitStr != " " {
				if _, err := strconv.Atoi(bmcSplitStr); err == nil {
					bmcPid = bmcSplitStr
					break
				}
			}
		}
		fmt.Println("The PID of BMC energy monitoring is " + bmcPid + ". Remember to kill if you exit unexpectedly")
		time.Sleep(3000 * time.Millisecond)
	}

	for i, socPort := range socPortList {
		wg.Add(1)
		strToExec := "adb -s " + socIp + ":" + socPort + " " + adbCmdStrList[i]
		adbLogPath := config.GetWorkPath() + "/adb_log/" + socPort
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
	if energy {
		sshCmd := exec.Command("bash", "-c", "ssh -p " + bmcPort + " root@" + socIp + " \"" + "kill -9 " + bmcPid + "\"")
		if err := sshCmd.Run(); err != nil {
			fmt.Println("kill energy monitor failed:", err)
		}
	}
	fmt.Println("finish exec command, the first of them is:\n" + adbCmdStrList[0])
}


