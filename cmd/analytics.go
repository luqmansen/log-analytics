package cmd

import (
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora"
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

func printLine(f *os.File, time time.Duration) {

	scanner := bufio.NewScanner(f)
	defer f.Close()

	for scanner.Scan() {
		func() {
			line := strings.Fields(scanner.Text())
			date, err := parseDate(parseLine(line))
			if err != nil {
				panic(err)
			}
			if checkTime(date, time) {
				if len(line) > 20 {
					fmt.Println(
						aurora.Green(line[1]).String() +
						" " + aurora.BrightBlue(line[2]).String() +
						" " + aurora.BrightYellow(line[3]+ line[4]).String()  +
						" " + aurora.Green(line[5]).String() +
						" " + aurora.BrightBlue(line[6]).String() +
						" " + aurora.BgBrightBlue(line[7]).String() +
						" " + aurora.BrightMagenta(line[8]).String() +
						" " + aurora.BrightRed(line[9]).String())
				}
				printCount++
			}
			lineCount++
		}()
	}
}

func isOpenNewFile() bool {
	if lineCount == printCount {
		lineCount, printCount = 0, 0
		return true
	}
	return false
}

func parseLine(s []string) string {
	if len(s) > 3 {
		return strings.Replace(s[3], "[", "", -1)
	} else {
		return ""
	}
}

func parseDate(s string) (time.Time, error) {
	return time.Parse(layout, s)
}

func checkTime(logTime time.Time, mins time.Duration) bool {

	if logTime.After(time.Now().Add(-mins * time.Minute)) {
		return true
	}
	return false
}

func analytics(cmd *cobra.Command, args []string) {

	dir, _ := cmd.Flags().GetString("directory")
	name, _ := cmd.Flags().GetString("file-name-pattern")
	mins, err := cmd.Flags().GetInt("time")

	if err != nil {
		fmt.Println(err)
	}
	number := 1
	for {
		f, err := os.Open(dir + "/" + name + strconv.Itoa(number))
		if os.IsNotExist(err) {
			break
		}
		printLine(f, time.Duration(mins))
		number++
		if !isOpenNewFile() {
			break
		}
	}
}