package main

import (
	"fmt"
	"github.com/donething/utils-go/dofile"
	"path/filepath"
	"strings"
)

const (
	backupDir = `D:/MyData/Setting/Windows`

	// git bash的etc配置文件
	// bashEtcPath  = `C:/Program Files/Git/etc/bash.bashrc`
	bashEtcPath  = `E:/Temp/bash.bashrc`
	bashTagBegin = `###### 自定义配置开始 ######`
	bashTagEnd   = `###### 自定义配置结束 ######`
	// git bash的界面配置文件
	bashUIPath = `C:/Users/Doneth/.minttyrc`
)

// 备份git bash的etc配置文件
func BackupBashEtc() (err error) {
	// 读取配置文件
	bs, err := dofile.Read(bashEtcPath)
	if err != nil {
		return
	}

	etc := string(bs)
	if !strings.Contains(etc, bashTagBegin) || !strings.Contains(etc, bashTagEnd) {
		return fmt.Errorf("没有找到需要备份的自定义配置")
	}

	// 提取需要备份的内容，并保存到目的文件
	text := etc[strings.Index(etc, bashTagBegin) : strings.Index(etc, bashTagEnd)+len(bashTagEnd)]
	dstPath := filepath.Join(backupDir, filepath.Base(bashEtcPath))
	_, err = dofile.Write([]byte(text), dstPath, dofile.WRITE_TRUNC, 0644)
	return err
}

// 恢复git bash的etc配置文件
func RestoreBashEtc() (err error) {
	// 去除现在的配置文件中的，已存在的自定义配置的内容
	bs, err := dofile.Read(bashEtcPath)
	if err != nil {
		return
	}
	etc := string(bs)
	if strings.Index(etc, bashTagBegin) >= 0 && strings.Index(etc, bashTagEnd) >= 0 {
		etc = etc[:strings.Index(etc, bashTagBegin)] + etc[strings.Index(etc, bashTagEnd)+len(bashTagEnd):]
	}

	// 读取备份配置
	src := filepath.Join(backupDir, filepath.Base(bashEtcPath))
	myEtc, err := dofile.Read(src)
	if err != nil {
		return
	}

	// 如果配置文件末尾没有换行符'\n'，则先追加
	if len(etc) > 0 && etc[len(etc)-1] != '\n' {
		etc = etc + "\n"
	}

	// 将备份配置内容追加到去除了自定义配置内容后，再写入当前配置文件中
	newEtc := etc + string(myEtc)
	_, err = dofile.Write([]byte(newEtc), bashEtcPath, dofile.WRITE_TRUNC, 0644)
	return
}

// 备份git bash的界面配置文件
func BackupBashUI() (err error) {
	dst := filepath.Join(backupDir, filepath.Base(bashUIPath))
	_, err = dofile.CopyFile(bashUIPath, dst, true)
	return err
}

// 恢复git bash的界面配置文件
func RestoreBashUI() (err error) {
	src := filepath.Join(backupDir, filepath.Base(bashUIPath))
	_, err = dofile.CopyFile(src, bashUIPath, true)
	return err
}
