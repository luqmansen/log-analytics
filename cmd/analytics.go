package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

func readFile(f *os.File, time int64) {
	scanner := bufio.NewScanner(f)
	defer f.Close()
	for scanner.Scan(){
		line := strings.Fields(scanner.Text())

		if checkTime(parseDate(parseLine(line)), time){
			fmt.Println(line[3] +" "+ line[4] + " " + line[5]+ " " + line[6]+ " " + line[7]+ " " + line[8]+ " " + line[9])
		}
	}
}

func parseLine(s []string) string{
	val := strings.Replace(s[3], "[", "", -1)
	//fmt.Println(val)
	return val
}

func parseDate(s string) time.Time{
	t, err := time.Parse(layout, s)
	checkEr(err)
	//fmt.Println(t.Unix())
	return t
}

func checkTime(logTime time.Time, mins int64) bool{

	if logTime.Unix()  > time.Now().Unix() -  (mins * 3600) {
		return true
	}

	return false
}


func analytics(cmd * cobra.Command, args []string){
	//dir, _ := cmd.Flags().GetString("directory")
	//fmt.Println(dir)
	f := openFile(".", "access.log.1")
	readFile(f, 100)



}
