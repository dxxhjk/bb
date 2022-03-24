package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

var configFileName = "/Users/shiboqing/GolandProjects/bb/config/config"
var config map[string]interface{}

func GetSocPortList() []string {
	socPortList := make([]string, 0)
	for _, port := range config["soc_port_list"].([]interface{}) {
		socPortList = append(socPortList, port.(string))
	}
	return socPortList
}

func GetSocIpListInternal() []string {
	socIpList := make([]string, 0)
	for _, ip := range config["soc_ip_list_internal"].([]interface{}) {
		socIpList = append(socIpList, ip.(string))
	}
	return socIpList
}

func GetSocPortInternal() string {
	return config["soc_port_internal"].(string)
}

func GetBmcIpInternal() string {
	return config["bmc_ip_internal"].(string)
}

func GetBmcPortInternal() string {
	return config["bmc_port_internal"].(string)
}

func GetWorkPath() string {
	return config["work_path"].(string)
}

func GetBaseIp() string {
	return config["base_ip"].(string)
}

func GetBmcPort() string {
	return config["bmc_port"].(string)
}

func GetLocalPort() string {
	return config["local_port"].(string)
}

func InitConfig() error {
	body, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}
	config = make(map[string]interface{})
	err = json.Unmarshal(body, &config)
	if err != nil {
		return err
	}
	if len(config["soc_port_list"].([]interface{})) != int(config["soc_num"].(float64)) ||
		len(config["soc_ip_list_internal"].([]interface{})) != int(config["soc_num"].(float64)){
		err = initSocPortList(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func initSocPortList(config map[string]interface{}) error {
	config["soc_port_list"] = []interface{}{}
	socBasePort, _ := strconv.Atoi(config["soc_base_port"].(string))
	for i := 1; i <= int(config["soc_num"].(float64)); i++ {
		socPort := strconv.Itoa(socBasePort + i)
		config["soc_port_list"] = append(config["soc_port_list"].([]interface{}), socPort)
	}
	config["soc_ip_list_internal"] = []interface{}{}
	socBaseIpInternal := config["soc_base_ip_internal"].(string)
	socBaseIpInternalLast := strings.Split(socBaseIpInternal, ".")[3]
	socBaseIpInternalPre := socBaseIpInternal[:len(socBaseIpInternal) - len(socBaseIpInternalLast)]
	socBaseIpInternalNum, _ := strconv.Atoi(socBaseIpInternalLast)
	for i := 1; i <= int(config["soc_num"].(float64)); i++ {
		socIpInternal := strconv.Itoa(socBaseIpInternalNum + i)
		config["soc_ip_list_internal"] = append(config["soc_ip_list_internal"].([]interface{}), 
			socBaseIpInternalPre + socIpInternal)
	}
	data, err := json.MarshalIndent(config, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(configFileName, data, 0777)
	if err != nil {
		return err
	}
	return nil
}