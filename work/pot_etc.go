package work

import (
	"github.com/donething/utils-go/dofile"
	"log"
	"path/filepath"
)

// 备份Potplayer配置
func BackupPotEtc() {
	// 文件路径为空白""，不需要备份
	if BashEtcPath[OS] == "" {
		return
	}

	dst := filepath.Join(Dir[OS], filepath.Base(PotEtcPath[OS]))
	_, err := dofile.CopyFile(PotEtcPath[OS], dst, true)
	if err != nil {
		log.Printf("备份Potplayer的配置文件出错：%s\n", err)
		return
	}
	log.Printf("备份Potplayer的配置文件完成：[%s] ==> [%s]\n", PotEtcPath[OS], dst)
}

// 恢复Potplayer配置
func RestorePotEtc() {
	// 文件路径为空白""，不需要备份
	if BashEtcPath[OS] == "" {
		return
	}

	src := filepath.Join(Dir[OS], filepath.Base(PotEtcPath[OS]))
	_, err := dofile.CopyFile(src, PotEtcPath[OS], true)
	if err != nil {
		log.Printf("恢复Potplayer的配置文件出错：%s\n", err)
		return
	}
	log.Printf("恢复Potplayer的配置文件完成：[%s] ==> [%s]\n", src, PotEtcPath[OS])
}
