package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
}
func alertNames(keyName []string, keyTime []string) []string {
	res := make(map[string][]string)
	for i := 0; i < len(keyName); i++ {
		if _, ok := res[keyName[i]]; !ok {
			res[keyName[i]] = make([]string, 0)
		}
		res[keyName[i]] = append(res[keyName[i]], keyTime[i])
	}

	r := make([]string, 0)
	for n, v := range res {
		if !check(v) {
			r = append(r, n)
		}
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}

func check(ts []string) bool {
	if len(ts) <= 2{
		return true
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i] < ts[j]
	})
	var ct, nt []string
	ct = strings.Split(ts[0], ":")
	ch, _ := strconv.ParseInt(ct[0], 10, 64)
	cm, _ := strconv.ParseInt(ct[1], 10, 64)
	flag := true
	for i := 0; i < len(ts)-2; i++ {
		flag = true
		nt = strings.Split(ts[i+1], ":")
		nh, _ := strconv.ParseInt(nt[0], 10, 64)
		nm, _ := strconv.ParseInt(nt[1], 10, 64)
		if nh-ch == 1 {
			if nm <= cm {
				flag = false
			}
		} else if nh-ch == 0 {
			flag = false
		}
		if !flag {
			nnt := strings.Split(ts[i+2], ":")
			nnh, _ := strconv.ParseInt(nnt[0], 10, 64)
			nnm, _ := strconv.ParseInt(nnt[1], 10, 64)
			if nnh-ch == 1 {
				if nnm <= cm {
					return false
				}
			} else if nnh-ch == 0 {
				return false
			}
		}
		ch = nh
		cm = nm
	}
	return true
}
