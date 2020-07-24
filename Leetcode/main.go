package main

import (
	"fmt"
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"path/filepath"
)

func main() {
	currPath, err := filepath.Abs("./Leetcode")
	if err != nil {
		fmt.Println(err)
		return
	}
	tools.GenTemplate(currPath, 1469, "find-all-the-lonely-nodes")
}
