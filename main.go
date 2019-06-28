// 备份与恢复
// 部分文件的备份和恢复，需要管理员权限
package main

import (
	"github.com/donething/utils-go/dofile"
	"github.com/donething/utils-go/dolog"
	"log"
	"os"
	"path"
	"path/filepath"
	"pc-backup-go/config"
	"strings"
	"time"
)

const prompt = "\n请输出正确参数：./pc-backup-go.exe config_path.json choice\n1. config_path.json表示配置文件路径，可以填\"\"（空字符串，而不是省略）\n2. choice表示要进行的操作 1：备份；2：还原"

var (
	// 日志文件
	lf *os.File
)

func init() {
	// 保存日志到文件
	var err error
	lf, err = dolog.LogToFile(dolog.LOG_NAME, dofile.WRITE_TRUNC, dolog.LOG_FORMAT)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("参数错误。%s\n", prompt)
	}
	// 初始化配置
	confPath := os.Args[1]
	if strings.TrimSpace(confPath) == "" {
		confPath = path.Join(".", config.DefaultName)
	}
	config.Init(confPath)

	switch os.Args[2] {
	case "1":
		log.Println("备份操作开始：")
		// 创建保存备份文件的目录
		err := os.MkdirAll(config.Dir, 0755)
		if err != nil {
			log.Fatalf("创建备份目录出错：%s\n", err)
		}

		// 备份
		backup()

		err = dofile.OpenAs(config.Dir)
		if err != nil {
			log.Fatalf("在资源管理器中显示备份目录时出错：%s\n", err)
		}
		log.Println("备份操作完成")
		defer lf.Close()

		// 如果立刻复制文件，可能导致上条log没有时间被写入文件，所以等待一会
		time.Sleep(2 * time.Second)

		_, err = dofile.CopyFile(dolog.LOG_NAME, filepath.Join(config.Dir, dolog.LOG_NAME), true)
		if err != nil {
			log.Printf("复制日志出错：%s\n", err)
		}
	case "2":
		log.Println("恢复操作开始：")
		restore()
		log.Println("恢复操作完成")
	default:
		log.Printf("参数错误。%s\n", prompt)
	}
}
