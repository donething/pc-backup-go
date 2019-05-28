package main

import (
	"github.com/donething/utils-go/dofile"
	"log"
)

// 复制文件
func copyFile(src string, dst string, tag string) {
	// 文件路径为空白""，不需要备份
	if src == "" {
		return
	}

	_, err := dofile.CopyFile(src, dst, true)
	if err != nil {
		log.Printf("复制%s出错：%s\n", tag, err)
		return
	}
	log.Printf("复制%s成功：[%s] ==> [%s]\n", tag, src, dst)
}
