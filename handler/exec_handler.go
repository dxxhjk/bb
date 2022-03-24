package handler

import (
	"bb/adb"
	"bb/config"
	"bb/util"
	"fmt"
)

func Exec(startSocPort, socNum, command string, energyMonitor bool) {
	socIp := config.GetBaseIp()
	socPortList := config.GetSocPortList()
	socPortList, err := util.GetDesignatedPortList(startSocPort, socNum, socPortList)
	if err != nil {
		fmt.Println(err)
		return
	}
	adb.Init(socIp, socPortList)
	adb.Shell(socIp, socPortList, command, energyMonitor)
}

func ExecInternal(startSocIp, socNum, command string, energyMonitor bool) {
	socIpList := config.GetSocIpListInternal()
	socPort := config.GetSocPortInternal()
	socIpList, err := util.GetDesignatedIpList(startSocIp, socNum, socIpList)
	if err != nil {
		fmt.Println(err)
		return
	}
	adb.InitInternal(socIpList, socPort)
	adb.ShellInternal(socIpList, config.GetBmcPortInternal(), socPort, command, energyMonitor)
}
