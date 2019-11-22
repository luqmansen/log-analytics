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

func TestParseLine(t *testing.T){
	s := []string{"a", "b", "c", "[stuff"}
	str := parseLine(s)
	if str != "stuff"{
		t.Errorf("parseLine([]string{\"a\", \"b\", \"[stuff\", \"d\"}) failed, expected %v, got %v", "stuff", str)
	}

	s = []string{"a", "b", "c"}
	str = parseLine(s)
	if str != ""{
		t.Errorf("parseLine([]string{\"a\", \"b\", \"c\"}) failed, expected %v, got %v", nil, str)
	}
}
