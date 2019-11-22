package cmd

import (
	"testing"
	"time"
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

func TestParseDate(t *testing.T){
	s := "02/Jan/2006:15:04:05"
	time2 := time.Date(2006, 1, 2, 15, 04, 05, 0, time.UTC)
	time1, err := parseDate(s)
	if time1 != time2{
		t.Errorf("parseDate(02/Jan/2006:15:04:05) failed, expected %v, got %v", time2, time1)
	}
	if err != nil{
		t.Errorf("parseDate(02/Jan/2006:15:04:05) failed, expected %v, got %v", nil, err)
	}
	_, err = parseDate("random string")
	if err == nil{
		t.Errorf("parseDate(\"random string\") failed, expected %v, got %v", nil, err)
	}
}
