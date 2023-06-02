package quick_sort

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func genSlice() ([]int, int) {
	res := make([]int, 0)
	sum := 0
	for i := 0; i < 100; i++ {
		res = append(res, i)
		sum += i
		if i%2 == 0 {
			res = append(res, i)
			sum += i
		}
	}
	return res, sum
}

func shuffleIntSlice(res []int) {
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
}

func randIntSlice() ([]int, int) {
	res := make([]int, 0)
	sum := 0
	for i := 0; i < 100; i++ {
		v := rand.Intn(100)
		res = append(res, v)
		sum += v
	}
	return res, sum
}

func checkSlice(src []int) (bool, int) {
	sum := src[0]
	flag := true
	for i := 1; i < len(src); i++ {
		if src[i] < src[i-1] {
			flag = false
		}
		sum += src[i]
	}
	return flag, sum
}

func sameSlice() []int {
	res := make([]int, 100)
	for i := 0; i < 100; i++ {
		res[i] = 100
	}
	return res
}

func TestQuickSort(t *testing.T) {
	PatchConvey("QuickSort", t, func() {
		for i := 0; i < 100; i++ {
			res, sum := genSlice()
			fmt.Println("original: ", res)
			shuffleIntSlice(res)
			fmt.Println("shuffled: ", res)
			quickSort(res)
			fmt.Println("sorted: ", res)
			result, resSum := checkSlice(res)
			So(result, ShouldBeTrue)
			So(resSum, ShouldEqual, sum)
		}
		for i := 0; i < 100; i++ {
			res, sum := randIntSlice()
			fmt.Println("original: ", res)
			quickSort(res)
			fmt.Println("sorted: ", res)
			result, resSum := checkSlice(res)
			So(result, ShouldBeTrue)
			So(resSum, ShouldEqual, sum)
		}

		res := sameSlice()
		fmt.Println("original: ", res)
		quickSort(res)
		fmt.Println("sorted: ", res)
		result, resSum := checkSlice(res)
		So(result, ShouldBeTrue)
		So(resSum, ShouldEqual, 100*100)
	})
}
