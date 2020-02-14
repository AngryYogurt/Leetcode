package tools

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

func AssertObjectEqual(expected, actual interface{}){
	if assert.ObjectsAreEqual(expected, actual){
		fmt.Println(fmt.Sprintf("Success. Expected: %+v Actual: %+v", expected, actual))
		return
	}
	fmt.Println(fmt.Sprintf("Failed. Expected: %+v Actual: %+v", expected, actual))
}

