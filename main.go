// 备份与恢复
// 部分文件的备份和恢复，需要管理员权限
package main

import (
	"github.com/donething/utils-go/dofile"
	"log"
	"os"
)

const prompt = `1：备份；2：还原`

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("参数错误。%s\n", prompt)
	}
	switch os.Args[1] {
	case "1":
		backup()
	case "2":
		restore()
	default:
		log.Printf("参数错误。%s\n", prompt)
	}
}

// 备份
func backup() {
	// 备份bash配置
	err := BackupBashEtc()
	if err != nil {
		log.Printf("备份bash配置文件出错：%s\n", err)
	} else {
		log.Printf("备份bash配置文件完成")
	}

	// 备份bash界面配置
	err = BackupBashUI()
	if err != nil {
		log.Printf("备份bash界面配置文件出错：%s\n", err)
	} else {
		log.Printf("备份bash界面配置文件完成")
	}

	err = dofile.OpenAs(backupDir)
	if err != nil {
		log.Printf("在资源管理器中显示备份目录时出错：%s\n", err)
	}
}

// 恢复
func restore() {
	err := RestoreBashEtc()
	if err != nil {
		log.Printf("恢复bash配置文件出错：%s\n", err)
	} else {
		log.Printf("恢复bash配置文件完成")
	}

	err = RestoreBashUI()
	if err != nil {
		log.Printf("恢复bash界面配置出错：%s\n", err)
	} else {
		log.Printf("恢复bash界面配置完成")
	}
}
