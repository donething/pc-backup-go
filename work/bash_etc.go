package work

import (
	"github.com/donething/utils-go/dofile"
	"log"
	"path/filepath"
	"strings"
)

const (
	// bash的etc配置文件中自定义部分的开始和结束标志
	bashTagBegin = `###### 自定义配置开始 ######`
	bashTagEnd   = `###### 自定义配置结束 ######`
)

// 备份git bash的etc配置文件
func BackupBashEtc() {
	// 文件路径为空白""，不需要备份
	if BashEtcPath[OS] == "" {
		return
	}

	// 读取配置文件
	bs, err := dofile.Read(BashEtcPath[OS])
	if err != nil {
		log.Printf("读取bash配置文件出错：%s\n", err)
		return
	}

	etc := string(bs)
	if !strings.Contains(etc, bashTagBegin) || !strings.Contains(etc, bashTagEnd) {
		log.Printf("没有找到需要备份的自定义配置\n")
		return
	}

	// 提取需要备份的内容，并保存到目的文件
	text := etc[strings.Index(etc, bashTagBegin) : strings.Index(etc, bashTagEnd)+len(bashTagEnd)]
	dstPath := filepath.Join(Dir[OS], filepath.Base(BashEtcPath[OS]))
	_, err = dofile.Write([]byte(text), dstPath, dofile.WRITE_TRUNC, 0644)
	if err != nil {
		log.Printf("备份bash配置出错：%s\n", err)
		return
	}
	log.Printf("备份bash配置文件成功：[%s] ==> [%s]\n", BashEtcPath[OS], dstPath)
}

// 恢复git bash的etc配置文件
func RestoreBashEtc() {
	// 文件路径为空白""，不需要备份
	if BashEtcPath[OS] == "" {
		return
	}

	// 去除现在的配置文件中的，已存在的自定义配置的内容
	bs, err := dofile.Read(BashEtcPath[OS])
	if err != nil {
		log.Printf("读取bash配置文件出错：%s\n", err)
		return
	}
	etc := string(bs)
	if strings.Index(etc, bashTagBegin) >= 0 && strings.Index(etc, bashTagEnd) >= 0 {
		etc = etc[:strings.Index(etc, bashTagBegin)] + etc[strings.Index(etc, bashTagEnd)+len(bashTagEnd):]
	}

	// 读取备份配置
	src := filepath.Join(Dir[OS], filepath.Base(BashEtcPath[OS]))
	myEtc, err := dofile.Read(src)
	if err != nil {
		log.Printf("读取bash的配置备份文件出错：%s\n", err)
		return
	}

	// 如果配置文件末尾没有换行符'\n'，则先追加
	if len(etc) > 0 && etc[len(etc)-1] != '\n' {
		etc = etc + "\n"
	}

	// 将备份配置内容追加到去除了自定义配置内容后，再写入当前配置文件中
	newEtc := etc + string(myEtc)
	_, err = dofile.Write([]byte(newEtc), BashEtcPath[OS], dofile.WRITE_TRUNC, 0644)
	if err != nil {
		log.Printf("恢复bash配置出错：%s\n", err)
		return
	}
	log.Printf("恢复bash配置文件成功：[%s] ==> [%s]\n", src, BashEtcPath[OS])
}
