// 备份与恢复
// 部分文件的备份和恢复，需要管理员权限
package main

import (
	"github.com/donething/utils-go/dofile"
	"github.com/donething/utils-go/dolog"
	"log"
	"os"
	"path/filepath"
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
		log.Println("备份操作开始：")
		// 创建保存备份文件的目录
		err := os.MkdirAll(Dir[OS], 0755)
		if err != nil {
			log.Fatalf("创建备份目录出错：%s\n", err)
		}

		// 备份
		backup()
		_, err = dofile.CopyFile(dolog.LOG_NAME, filepath.Join(Dir[OS], dolog.LOG_NAME), true)
		if err != nil {
			log.Printf("复制日志出错：%s\n", err)
		}

		err = dofile.OpenAs(Dir[OS])
		if err != nil {
			log.Fatalf("在资源管理器中显示备份目录时出错：%s\n", err)
		}
		log.Println("备份操作完成")
	case "2":
		log.Println("恢复操作开始：")
		restore()
		log.Println("恢复操作完成")
	default:
		log.Printf("参数错误。%s\n", prompt)
	}
}
