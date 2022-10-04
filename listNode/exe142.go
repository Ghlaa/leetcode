package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hashmap := map[*ListNode]bool{}
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	temp := head
	for temp != nil {
		if hashmap[temp] == false {
			hashmap[temp] = true
			temp = temp.Next

		} else {
			return temp.Next
		}

	}
	return nil
}
