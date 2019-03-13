package main

import (
	"testing"
)

func TestBackupBashEtc(t *testing.T) {
	err := BackupBashEtc()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBackupBashUI(t *testing.T) {
	err := BackupBashUI()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRestoreBashEtc(t *testing.T) {
	err := RestoreBashEtc()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRestoreBashUI(t *testing.T) {
	err := RestoreBashUI()
	if err != nil {
		t.Fatal(err)
	}
}
