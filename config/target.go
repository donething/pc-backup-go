package config

import (
	"os"
	"path/filepath"
	"runtime"
)

const (
	OSWin   = "windows"
	OSLinux = "linux"
)

var (
	OS          = runtime.GOOS            // 当前操作系统
	Dir         = ""                      // 备份目录
	BashEtcPath = make(map[string]string) // bash etc配置
	BashUIPath  = make(map[string]string) // bash UI配置
	PotEtcPath  = make(map[string]string) // potplayer 配置
)

// 初始化配置
func Init(confPath string) {
	initConf(confPath)
	// 备份目录
	Dir = Conf.Dir

	// bash etc配置
	BashEtcPath[OSWin] = filepath.Join(os.Getenv("USERPROFILE"), ".bashrc")
	BashEtcPath[OSLinux] = filepath.Join(os.Getenv("HOME"), ".bashrc")

	// bash UI配置
	BashUIPath[OSWin] = filepath.Join(os.Getenv("USERPROFILE"), ".minttyrc")

	// potplayer 配置
	PotEtcPath[OSWin] = filepath.Join(os.Getenv("APPDATA"), "PotPlayerMini64/PotPlayerMini64.ini")
}
