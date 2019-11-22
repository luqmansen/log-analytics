package cmd

import (
	"testing"
)

func TestIsOpenNewFile(t *testing.T) {
	lineCount = 10
	printCount = 10

	b := isOpenNewFile()

	if b == false {
		t.Errorf("isOpenNewFile() failed, expected %v, got %v", true, b)
	}
	printCount++

	b = isOpenNewFile()
	if b == true {
		t.Errorf("isOpenNewFile() failed, expected %v, got %v", false, b)
	}

}
