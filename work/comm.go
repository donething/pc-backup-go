package work

import "runtime"

const (
	OS_WIN   = "windows"
	OS_LINUX = "linux"
)

var (
	OS          = runtime.GOOS            // 当前操作系统
	Dir         = make(map[string]string) // 备份目录
	BashEtcPath = make(map[string]string) // bash etc配置
	BashUIPath  = make(map[string]string) // bash UI配置
)

// 初始化配置
func init() {
	// 备份目录
	Dir[OS_WIN] = "D:/MyData/Setting/Windows/pc-backup-go"
	Dir[OS_LINUX] = "/home/doneth/MyData/Setting/Windows/pc-backup-go"

	// bash etc配置
	BashEtcPath[OS_WIN] = "C:/Program Files/Git/etc/bash.bashrc"
	BashEtcPath[OS_LINUX] = "/home/doneth/.bashrc"

	// bash UI配置
	BashUIPath[OS_WIN] = "C:/Users/Doneth/.minttyrc"
	BashUIPath[OS_LINUX] = ""

}
