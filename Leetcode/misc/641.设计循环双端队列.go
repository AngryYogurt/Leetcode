/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	head   int
	tail   int
	size   int
	limit  int
	values []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		head:   0,
		tail:   0,
		size:   0,
		limit:  k,
		values: make([]int, k+1),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head--
	if this.head < 1 {
		this.head = len(this.values) - 1
	}
	this.values[this.head] = value
	this.size++
	if this.tail == 0 {
		this.tail = this.head
	}
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.tail++
	if this.tail > len(this.values)-1 {
		this.tail = 1
	}
	this.values[this.tail] = value
	this.size++
	if this.head == 0 {
		this.head = this.tail
	}
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head++
	if this.head >= len(this.values) {
		this.head = 1
	}
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail--
	if this.tail < 1 {
		this.tail = len(this.values) - 1
	}
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.values[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size <= 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size >= this.limit
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

