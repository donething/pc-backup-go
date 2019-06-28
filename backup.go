package main

import (
	"github.com/donething/utils-go/dofile"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"log"
	"path/filepath"
	. "pc-backup-go/config"
	"strings"
)

// 备份
func backup() {
	// bashrc
	copyFile(BashEtcPath[OS], filepath.Join(Dir, filepath.Base(BashEtcPath[OS])), "[bashrc]")
	// bash ui
	copyFile(BashUIPath[OS], filepath.Join(Dir, filepath.Base(BashUIPath[OS])), "[bashUI]")
	// potplayer配置
	copyFile(PotEtcPath[OS], filepath.Join(Dir, filepath.Base(PotEtcPath[OS])), "[Pot.ini]")
	// backupPot()
}

// 备份Potplayer
// 由于配置文件时UTF16LE格式，需要先转换为UTF8，截取指定内容后，再转为UTF16LE后保存
// *** 还未完成，中文部分乱码
func backupPot() {
	bs, err := dofile.Read(PotEtcPath[OS])
	if err != nil {
		log.Printf("读取Potplayer的配置文件出错：%s", err)
		return
	}

	result, _, err := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewEncoder(), bs)
	if err != nil {
		log.Printf("将Potplayer的配置由UTF16LE转为UTF8时出错：%s", err)
		return
	}
	confText := string(result)
	log.Println(confText)
	newConf1 := confText[0:strings.Index(confText, `[RememberFiles]`)]
	newConf2 := confText[strings.Index(confText, `[Settings]`):]
	newConf := newConf1 + newConf2
	_, err = dofile.Write([]byte(newConf), filepath.Join(Dir, filepath.Base(PotEtcPath[OS])), dofile.WRITE_TRUNC, 0644)
	if err != nil {
		log.Printf("备份Potplayer的配置文件出错：%s", err)
		return
	}
}

// 恢复
func restore() {
	// bashrc
	copyFile(filepath.Join(Dir, filepath.Base(BashEtcPath[OS])), BashEtcPath[OS], "[bashrc]")
	// bash ui
	copyFile(filepath.Join(Dir, filepath.Base(BashUIPath[OS])), BashUIPath[OS], "[bashUI]")
	// potplayer配置
	copyFile(filepath.Join(Dir, filepath.Base(PotEtcPath[OS])), PotEtcPath[OS], "[Pot.ini]")
}
