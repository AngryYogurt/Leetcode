/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	max, l := 1, len(s)
	res := s[:1]
	for idx, _ := range s {
		if l-idx <= max/2 {
			break
		}
		length, odd := check(s, idx)
		if length > max {
			max = length
			if odd {
				res = s[idx-length/2 : idx+length/2+1]
			} else {
				res = s[idx-length/2+1 : idx+length/2+1]
			}
		}
	}

	return res
}

func check(s string, idx int) (int, bool) {
	// even
	even := 0
	left, right := idx, idx+1
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		even++
		left--
		right++
	}
	odd := 0
	// odd
	left, right = idx-1, idx+1
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		odd++
		left--
		right++
	}
	if even > odd {
		return even * 2, false
	}
	return odd*2 + 1, true
}

// @lc code=end

