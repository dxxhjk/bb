package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

var configFileName = "./config/config"
var config map[string]interface{}

func GetSocIpList() []string {
	socIpList := make([]string, 0)
	for _, ip := range config["soc_ip_list"].([]interface{}) {
		socIpList = append(socIpList, ip.(string))
	}
	return socIpList
}

func GetWorPath() string {
	return config["work_path"].(string)
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
	if len(config["soc_ip_list"].([]interface{})) != int(config["soc_num"].(float64)) {
		err = initSocIpList(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func initSocIpList(config map[string]interface{}) error {
	config["soc_ip_list"] = []interface{}{}
	for i := 1; i <= int(config["soc_num"].(float64)); i++ {
		socIp := strconv.Itoa(i)
		if len(socIp) == 1 {
			socIp = config["soc_base_ip"].(string) + "0" + socIp
		} else {
			socIp = config["soc_base_ip"].(string) + socIp
		}
		config["soc_ip_list"] = append(config["soc_ip_list"].([]interface{}), socIp)
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