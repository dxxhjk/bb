package handler

import (
	"bb/adb"
	"bb/config"
	"bb/util"
	"fmt"
)

func DistributeFile(startSocPort, socNum, fileToDistribute, destination string) {
	socIp := config.GetBaseIp()
	socPortList := config.GetSocPortList()
	socPortList, err := util.GetDesignatedPortList(startSocPort, socNum, socPortList)
	if err != nil {
		fmt.Println(err)
		return
	}
	adb.Init(socIp, socPortList)
	adb.Push(socIp, socPortList, fileToDistribute, destination)
}

func DistributeFileInternal(startSocIp, socNum, fileToDistribute, destination string) {
	socIpList := config.GetSocIpListInternal()
	socPort := config.GetSocPortInternal()
	socIpList, err := util.GetDesignatedIpList(startSocIp, socNum, socIpList)
	if err != nil {
		fmt.Println(err)
		return
	}
	adb.InitInternal(socIpList, socPort)
	adb.PushInternal(socIpList, socPort, fileToDistribute, destination)
}