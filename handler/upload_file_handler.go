package handler

import (
	"bb/config"
	"fmt"
)

func UploadFile(filePath, loginName string) {
	localPath := config.GetWorkPath() + "/file/"
	localIp := config.GetBaseIp()
	localPort := config.GetLocalPort()
	fmt.Println("use command:")
	fmt.Println("scp -r -P " + localPort + " " +
		filePath + " " + loginName + "@" +
		localIp + ":" + localPath)
}