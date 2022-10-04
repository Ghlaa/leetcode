package main

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}
type MyLinkedList struct {
	dummy *Node
}

func Constructor() MyLinkedList {
	node := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	node.Pre = node
	node.Next = node
	return MyLinkedList{node}
}

func (this *MyLinkedList) Get(index int) int {
	node := this.dummy.Next
	if node == this.dummy {
		return -1 //表示当前数组还是没有值
	}
	for node != this.dummy && index > 0 {
		index--
		node = node.Next
	}
	if index != 0 {
		return -1 //索引无效
	}
	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Pre:  dummy,
		Next: dummy.Next,
	}
	dummy.Next.Pre = node
	dummy.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy,
		Pre:  dummy.Pre,
	}
	dummy.Pre.Next = node
	dummy.Pre = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := this.dummy.Next
	for node != this.dummy && index > 0 {
		index--
		node = node.Next
	}
	if index != 0 {
		return
	}
	temp := &Node{
		Val:  val,
		Next: node,
		Pre:  node.Pre,
	}
	node.Pre.Next = temp
	node.Pre = temp

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	head := this.dummy.Next
	if head == this.dummy {
		return //说明此时链表为空
	}
	for head != this.dummy && index > 0 {
		index--
		head = head.Next
	}
	if index != 0 {
		return
	}
	if head == this.dummy {
		return
	}
	head.Pre.Next = head.Next
	head.Next.Pre = head.Pre
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
