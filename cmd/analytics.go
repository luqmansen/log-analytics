package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)


func checkEr(e error)  {
	if e != nil{
		panic(e)
	}
}

func checkFile(dir string, filename string) (*os.File, error){
	dat, err := ioutil.ReadFile(dir + "/" + filename)

	checkEr(err)
	fmt.Println(string(dat))
	fmt.Println(dir + "/"  +filename)
	f, err := os.Open(dir + "/"  +filename)
	checkEr(err)
	return f, nil
}


func analytics(cmd * cobra.Command, args []string){
	dir, _ := cmd.Flags().GetString("directory")
	fmt.Println(dir)
	checkFile(dir, "access.log.1")


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
