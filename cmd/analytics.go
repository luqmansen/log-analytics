package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	layout = "02/Jan/2006:15:04:05"
)

func checkEr(e error)  {
	if e != nil{
		panic(e)
	}
}

func openFile(dir string, filename string) (*os.File){
	file, err := os.Open(dir + "/" + filename)
	fmt.Println(dir +"/"+ filename)
	checkEr(err)
	return file
}

func readFile(f *os.File) *bufio.Scanner{
	scanner := bufio.NewScanner(f)
	defer f.Close()
	for scanner.Scan(){
		line := strings.Fields(scanner.Text())
		parseLine(line)
	}
	return scanner
}

func parseLine(s []string) bool{
	val := strings.Replace(s[3], "[", "", -1)
	//fmt.Println(val)
	t, err := time.Parse(layout, val)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())

	return true
}


func analytics(cmd * cobra.Command, args []string){
	//dir, _ := cmd.Flags().GetString("directory")
	//fmt.Println(dir)
	f := openFile(".", "access.log.1")
	readFile(f)



}

func visit(files *[]string) filepath.WalkFunc{
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}
