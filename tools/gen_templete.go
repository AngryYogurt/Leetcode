package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenTemplate(sn int64, problem string) {
	name := strings.ReplaceAll(problem, "-", "_")
	dir := fmt.Sprintf("%d_%s", sn, name)
	if _, err := os.Stat(dir); os.IsExist(err) {
		fmt.Println(fmt.Sprintf("%s is already exist", dir))
	}
	os.Mkdir(dir, os.ModeDir)
	f, _ := os.Create(filepath.Join(dir, name+".go"))
	f.WriteString(`package main

func main(){
	
}
`)
}
