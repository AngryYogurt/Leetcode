package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(thousandSeparator(-12343242))
	fmt.Println(thousandSeparator(987))

}

func thousandSeparator(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		n = -n
		neg = true
	}
	ns := strconv.FormatInt(int64(n), 10)
	t := 0
	for i := len(ns) - 1; i >= 0; i-- {
		res = string(ns[i]) + res
		t++
		if t == 3 {
			t = 0
			res = "." + res
		}
	}
	res = strings.Trim(res, ".")
	if neg {
		return "-" + res
	}
	return res
}
