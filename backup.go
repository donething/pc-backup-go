package main

import (
	"path/filepath"
	. "pc-backup-go/config"
)

// 备份
func backup() {
	// bashrc
	copyFile(BashEtcPath[OS], filepath.Join(Dir, filepath.Base(BashEtcPath[OS])), "[bashrc]")
	// bash ui
	copyFile(BashUIPath[OS], filepath.Join(Dir, filepath.Base(BashUIPath[OS])), "[bashUI]")
	// potplayer配置
	copyFile(PotEtcPath[OS], filepath.Join(Dir, filepath.Base(PotEtcPath[OS])), "[Pot.ini]")
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
