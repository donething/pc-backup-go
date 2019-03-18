// 备份与恢复
// 部分文件的备份和恢复，需要管理员权限
package main

import (
	"github.com/donething/utils-go/dofile"
	"github.com/donething/utils-go/dolog"
	"log"
	"os"
	"pc-backup-go/work"
)

const prompt = `1：备份；2：还原`

func init() {
	// 保存日志到文件
	err := dolog.Log2File(dolog.LOG_NAME, dofile.WRITE_TRUNC, dolog.LOG_FORMAT)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("参数错误。%s\n", prompt)
	}
	switch os.Args[1] {
	case "1":
		bk()
	case "2":
		rt()
	default:
		log.Printf("参数错误。%s\n", prompt)
	}
}

// 备份
func bk() {
	// 创建保存备份文件的目录
	err := os.MkdirAll(work.Dir[work.OS], 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 备份bash配置
	work.BackupBashEtc()
	// 备份bash界面配置
	work.BackupBashUI()

	err = dofile.OpenAs(work.Dir[work.OS])
	if err != nil {
		log.Printf("在资源管理器中显示备份目录时出错：%s\n", err)
	}
}

// 恢复
func rt() {
	work.RestoreBashEtc()
	work.RestoreBashUI()
}
