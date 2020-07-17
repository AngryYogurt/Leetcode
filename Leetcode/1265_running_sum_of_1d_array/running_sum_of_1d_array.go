package main

func main() {

}

/*   Below is the interface for ImmutableListNode, which is already defined for you.
 */
type ImmutableListNode struct {
}

func (this *ImmutableListNode) getNext() *ImmutableListNode {
	// return the next node.
	return nil
}

func (this *ImmutableListNode) printValue() {
	// print the value of this node.
}

/*
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：2.4 MB, 在所有 Go 提交中击败了100.00%的用户
 */

func printLinkedListInReverse(head *ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	head.printValue()
}
