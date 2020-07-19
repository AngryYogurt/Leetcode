package main

import (
	"fmt"
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"path/filepath"
)

func main() {
	currPath, err := filepath.Abs("./HackerRank")
	if err != nil {
		fmt.Println(err)
	}
	tools.GenTemplate(currPath, -1, "digit-products")
}
