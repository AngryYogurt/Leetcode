package main

import (
	"fmt"
	"github.com/AngryYogurt/leetcode/tools"
	"path/filepath"
)

func main() {
	currPath, err := filepath.Abs("./Leetcode")
	if err != nil {
		fmt.Println(err)
	}
	tools.GenTemplate(currPath, 5465, "number-of-nodes-in-the-sub-tree-with-the-same-label")
}
