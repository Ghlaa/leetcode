package main

func swapPairs(head *ListNode) *ListNode {
	pre := head
	if pre == nil { //如果链表为空的话
		return head
	}
	cur := pre.Next
	if cur == nil { //如果链表中元素只有一个的话
		return head
	}
	temp := cur.Next //先给temp随便赋值
	head = head.Next
	flag := false
	for flag == false {
		//temp=cur.Next
		if temp == nil { //元素个数为偶数的时候,并且已经处理到了最后的结果
			pre.Next = nil
			cur.Next = pre
			flag = true
		} else if temp.Next == nil { //当元素个数为奇数的时候，并且已经处理到了最后一个元素，temp就是最后一个元素了,temp.next就是指向了nil

			cur.Next = pre
			pre.Next = temp
			flag = true
			/*pre=temp
			cur=temp.Next*/
		} else if temp.Next != nil { //证明后面至少还有两个元素，可以继续转换链表
			cur.Next = pre
			pre.Next = temp.Next

			pre = temp
			cur = temp.Next
			temp = cur.Next
		}
	}
	return head
}
