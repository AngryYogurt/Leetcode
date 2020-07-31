package main

import (
	"fmt"
	"math/rand"
	"time"
)

// FIXME 不会写
func main() {
	rand.Seed(time.Now().Unix())
	res := int(rand.Int31n(10))
	src := make([]int, 0)
	for i := 0; i < 10; i++ {
		if i != res {
			for j := 0; j < 4; j++ {
				src = append(src, i)
			}
		}
		src = append(src, i)
	}
	rand.Shuffle(len(src), func(i, j int) {
		src[i], src[j] = src[j], src[i]
	})
	find(src)
	fmt.Println(src, len(src))
	fmt.Println(res)
}

func find(src []int) {
	v1, v2, v3, v4, v5 := 0, 0, 0, 0, 0
	for _, v := range src {
		v1 = (^v5) & (^v4) & (^v3) & (^v2) & (v1 ^ v)
		v2 = (^v5) & (^v5) & (^v5) & (^v5) &  (v2 ^ v)
		v3 = (^v5) & (^v5) & (^v5) & (^v5) &  (v3 ^ v)
		v4 = (^v5) & (^v5) & (^v5) & (^v5) &  (v4 ^ v)
		v5 = (^v5) & (^v5) & (^v5) & (^v5) &  (v5 ^ v)
	}
	fmt.Println(v1, v2, v3, v4, v5)
}
