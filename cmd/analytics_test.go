package cmd

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T){


	f , err := openFile("../", "access.log.1")
	//provide real file
	if err != nil{
		t.Errorf("openFile(\"., access.log.1\") failed, expected %v, got %v", nil, err)
	} else {
		t.Log(f)
	}

	//provide fake file
	_ , err2 := openFile(".", "access.log.x")
	if !os.IsNotExist(err2){
		t.Errorf("openFile(\".\", \"access.log.x\") failed, expected %v, got %v", nil, err2)
	}
}