package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 配置文件默认的名字
const DefaultName = `pc-backup-go.json`

type config struct {
	// 备份目录
	Dir string `json:"dir"`
}

var Conf = config{""}

func initConf(confPath string) {
	saveConfig(confPath + ".bak")
	parseConfig(confPath)
	if strings.TrimSpace(Conf.Dir) == "" {
		log.Fatalf("还未填写配置文件")
	}
}

// 解析配置文件
func parseConfig(path string) {
	// 如果配置文件不存在，创建
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("配置文件（%s）不存在", path)
		return
	}

	// 读取json配置文件，实例到Conf对象中
	data, errRead := ioutil.ReadFile(path)
	if errRead != nil {
		log.Printf("读取配置文件(%s)失败：%v", path, errRead)
		return
	}
	errParse := json.Unmarshal(data, &Conf)
	if errParse != nil {
		log.Printf("解析配置文件(%s)错误：%v", path, errParse)
	}
}

// 创建配置文件
func saveConfig(path string) bool {
	data, errJson := json.MarshalIndent(Conf, "", "\t")
	if errJson != nil {
		log.Printf("将结构体配置数据(%+v)转为json格式数据失败：%#v", Conf, errJson)
		return false
	}
	errWrite := ioutil.WriteFile(path, data, 0644)
	if errWrite != nil {
		log.Printf("保存配置信息到文件失败：%#v", errWrite)
		return false
	}
	log.Printf("配置(%s)保存完成", path)
	return true
}
