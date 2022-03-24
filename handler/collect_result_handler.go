package handler

import (
	"bb/adb"
	"bb/config"
	"bb/util"
	"fmt"
)

func CollectResult(fileToCollect string, localPathToSaveFiles string, startSocPort string, socNum string) {
	socIp := config.GetBaseIp()
	socPortList := config.GetSocPortList()
	socPortList, err := util.GetDesignatedPortList(startSocPort, socNum, socPortList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(localPathToSaveFiles)
	adb.Init(socIp, socPortList)
	adb.Pull(socIp, socPortList, fileToCollect, localPathToSaveFiles)
}

func CollectResultInternal(fileToCollect string, localPathToSaveFiles string, startSocIp string, socNum string) {
	socIpList := config.GetSocIpListInternal()
	socPort := config.GetSocPortInternal()
	socIpList, err := util.GetDesignatedIpList(startSocIp, socNum, socIpList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(localPathToSaveFiles)
	adb.InitInternal(socIpList, socPort)
	adb.PullInternal(socIpList, socPort, fileToCollect, localPathToSaveFiles)
}
