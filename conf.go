package main

import (
	"os"
	"path/filepath"
	"runtime"
)

const (
	OS_WIN   = "windows"
	OS_LINUX = "linux"
)

var (
	OS          = runtime.GOOS            // 当前操作系统
	Dir         = make(map[string]string) // 备份目录
	BashEtcPath = make(map[string]string) // bash etc配置
	BashUIPath  = make(map[string]string) // bash UI配置
	PotEtcPath  = make(map[string]string) // potplayer 配置
)

// 初始化配置
func init() {
	// 备份目录
	Dir[OS_WIN] = "D:/MyData/Setting/Windows/pc-backup-go"
	Dir[OS_LINUX] = filepath.Join(os.Getenv("HOME"), "MyData/Setting/Windows/pc-backup-go")

	// bash etc配置
	BashEtcPath[OS_WIN] = filepath.Join(os.Getenv("USERPROFILE"), ".bashrc")
	BashEtcPath[OS_LINUX] = filepath.Join(os.Getenv("HOME"), ".bashrc")

	// bash UI配置
	BashUIPath[OS_WIN] = filepath.Join(os.Getenv("USERPROFILE"), ".minttyrc")
	BashUIPath[OS_LINUX] = ""

	// potplayer 配置
	PotEtcPath[OS_WIN] = filepath.Join(os.Getenv("APPDATA"), "PotPlayerMini64/PotPlayerMini64.ini")
	PotEtcPath[OS_LINUX] = ""
}
