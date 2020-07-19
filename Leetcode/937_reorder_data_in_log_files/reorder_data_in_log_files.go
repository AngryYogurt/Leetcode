package main

import (
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"sort"
	"strings"
)

func main() {
	tools.AssertObjectEqual([]string{
		"b aq cojj", "5 ba iedrz", "8 fj dzz k", "63 gu psub", "bb wsrd kp",
		"5r 446 6 3", "6s 87979 5", "3r 587 01", "jc 3480612", "r5 6316 71",
	}, reorderLogFiles([]string{
		"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5",
		"3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71",
	}))
}

type L struct {
	Data string
	Id   string
	Log  string
}

type Ls []*L

func isDigit(a byte) bool {
	return a >= '0' && a <= '9'
}

func (l Ls) Less(i, j int) bool {
	if isDigit(l[i].Log[0]) {
		if isDigit(l[j].Log[0]) {
			return i < j
		}
		return false
	}
	if isDigit(l[j].Log[0]) {
		return true
	}
	if l[i].Log < l[j].Log {
		return true
	} else if l[i].Log > l[j].Log {
		return false
	} else {
		if l[i].Id < l[j].Id {
			return true
		} else {
			return false
		}
	}
}

func (l Ls) Len() int {
	return len(l)
}

func (l Ls) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func reorderLogFiles(logs []string) []string {
	l := make(Ls, 0)
	for _, log := range logs {
		s := strings.SplitN(log, " ", 2)
		l = append(l, &L{
			Data: log,
			Id:   s[0],
			Log:  s[1],
		})
	}
	sort.Sort(l)
	result := make([]string, len(l))
	for i, d := range l {
		result[i] = d.Data
	}
	return result
}
