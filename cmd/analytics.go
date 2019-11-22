package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	layout = "02/Jan/2006:15:04:05"
)

var lineCount, printCount int

func checkEr(e error) {
	if e != nil{
		panic(e)
	}
}

func printLine(f *os.File, time time.Duration) {

	scanner := bufio.NewScanner(f)
	defer f.Close()

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if checkTime(parseDate(parseLine(line)), time) {
			if len(line)>20{
				fmt.Println(line[1] + " " + line[2] + " " + line[3] + line[4] + " " + line[5] + " " + line[6] + " " + line[7] + " " + line[8] + " " + line[9])
			}
			printCount++
		}
		lineCount++
	}
}

func isOpenNewFile() bool {
	//fmt.Println(lineCount, printCount)
	if lineCount == printCount {
		lineCount, printCount = 0, 0
		return true
	}
	return false
}

func parseLine(s []string) string {
	var val string
	if len(s) > 3{
		val = strings.Replace(s[3], "[", "", -1)
	}
	return val
}

func parseDate(s string) time.Time {
	t, err := time.Parse(layout, s)
	checkEr(err)
	return t
}

func checkTime(logTime time.Time, mins time.Duration) bool {

	if logTime.After(time.Now().Add(- mins* time.Minute )){
		return true
	}
	return false
}

func analytics(cmd *cobra.Command, args []string) {

	dir, _ := cmd.Flags().GetString("directory")
	mins, err := cmd.Flags().GetInt("time")
	//fmt.Println(err)
	//fmt.Println("directory : " + dir)
	//fmt.Printf("time : %v \n ",time.Duration(mins)*time.Minute)

	//log := time.Now().Add(time.Duration(-mins) * time.Minute)

	//fmt.Printf("log %v",  log)

	if err != nil {
		fmt.Println(err)
	}
	number := 1
	for {
		f, err := os.Open(dir +"/"+ "access.log."+strconv.Itoa(number))
		if os.IsNotExist(err){
			//fmt.Println(err)
			break
		}
		printLine(f, time.Duration(mins))
		number++
		if !isOpenNewFile() {
			break
		}
	}

}
