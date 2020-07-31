package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenTemplate(path string, sn int64, problem string) {
	name := strings.ReplaceAll(problem, "-", "_")
	dir := name
	if sn >= 0 {
		dir = fmt.Sprintf("%d_%s", sn, name)
	}
	dir = filepath.Join(path, dir)
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("%s is already exist", dir))
		return
	}
	os.Mkdir(dir, os.ModeDir)
	f, _ := os.Create(filepath.Join(dir, name+".go"))
	f.WriteString(`package main

func main(){
	
}
`)
}
