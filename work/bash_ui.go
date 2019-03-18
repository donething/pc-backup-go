package work

import (
	"github.com/donething/utils-go/dofile"
	"log"
	"path/filepath"
)

// 备份git bash的界面配置文件
func BackupBashUI() {
	// 文件路径为空白""，不需要备份
	if BashUIPath[OS] == "" {
		return
	}

	dst := filepath.Join(Dir[OS], filepath.Base(BashUIPath[OS]))
	_, err := dofile.CopyFile(BashUIPath[OS], dst, true)
	if err != nil {
		log.Printf("备份bash的外观配置文件出错：%s\n", err)
		return
	}
	log.Printf("备份bash的外观配置文件完成：[%s] ==> [%s]\n", BashUIPath[OS], dst)
}

// 恢复git bash的界面配置文件
func RestoreBashUI() {
	// 文件路径为空白""，不需要备份
	if BashUIPath[OS] == "" {
		return
	}

	src := filepath.Join(Dir[OS], filepath.Base(BashUIPath[OS]))
	_, err := dofile.CopyFile(src, BashUIPath[OS], true)
	if err != nil {
		log.Printf("恢复bash的外观配置文件出错：%s\n", err)
		return
	}
	log.Printf("恢复bash的外观配置文件完成：[%s] ==> [%s]\n", src, BashUIPath[OS])

}
