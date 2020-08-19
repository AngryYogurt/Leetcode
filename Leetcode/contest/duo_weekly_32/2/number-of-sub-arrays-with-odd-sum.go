package main

import "fmt"

func main() {
	fmt.Println(canConvertString("jicfxaa", "ocxltbp", 15))
	fmt.Println(canConvertString("input", "ouput", 9))
	fmt.Println(canConvertString("aab", "bbb", 27))
}

func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	t1 := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		if t[i] == s[i] {
			continue
		}
		d := int(t[i]) - int(s[i])
		if d < 0 {
			d += 26
		}
		if d > k {
			return false
		}
		if _, ok := t1[d]; !ok {
			t1[d] = true
		} else {
			ok := false
			for d+26 <= k {
				d += 26
				if _, okj := t1[d]; !okj {
					ok = true
					t1[d] = true
					break
				}
			}
			if !ok {
				return false
			}
		}
	}
	return true
}
