package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
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
	if len(config["soc_port_list"].([]interface{})) != int(config["soc_num"].(float64)) {
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