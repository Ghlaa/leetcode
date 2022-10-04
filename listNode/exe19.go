package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	temp := dummy.Next
	if temp == nil {
		//头节点为空，直接返回
		return head
	}
	cnt := 1
	for temp.Next != nil {
		cnt += 1
		temp = temp.Next
	} //循环结束后temp指向最后一个节点,并且cnt表示节点的个数
	count := cnt + 1 - n //删除到处第n个元素，就是删除正数第count个
	if count < 1 {       //肯定有错误，比如总共只有五个元素，你说要删除倒数第八个
		return nil
	}
	temp1 := dummy
	if count == 1 && cnt == 1 { //总共只有一个元素，并且删除第一个的时候
		return nil
	}
	for temp1.Next != nil && count > 1 {
		temp1 = temp1.Next
		count--
	}
	temp1.Next = temp1.Next.Next
	return dummy.Next

}
