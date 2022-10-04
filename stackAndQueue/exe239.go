package stackAndQueue

type myQueue struct {
	queue []int
}

func NewMyQueue() *myQueue {
	return &myQueue{queue: []int{}}
}
func (m *myQueue) Front() int { //弹出队首元素，也就是最大值
	return m.queue[0]
}
func (m *myQueue) Push(a int) {
	for len(m.queue) != 0 {
		if m.queue[len(m.queue)-1] < a { //因为要对队列的元素进行排序，所以把要比a小的元素全部移除，把a放置对应的位置
			m.queue = m.queue[0 : len(m.queue)-1]
		} else {
			break
		}
	}
	m.queue = append(m.queue, a)
}
func (m *myQueue) Pop(a int) {
	if m.Front() == a { //如果需要移除的元素正好是当前的最大值，那移除，否则什么都不做
		m.queue = m.queue[1:len(m.queue)]
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	myQueue := NewMyQueue()
	for i := 0; i < k; i++ {
		myQueue.Push(nums[i])
	}
	res = append(res, myQueue.Front())
	myQueue.Pop(nums[0])
	for i := k; i < len(nums); i++ {
		myQueue.Push(nums[i])
		res = append(res, myQueue.Front())
		myQueue.Pop(nums[i-k+1])
	}
	return res
}
