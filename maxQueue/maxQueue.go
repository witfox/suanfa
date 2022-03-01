package maxqueue

import "container/list"

//队列的最大值
//请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
//若队列为空，pop_front 和 max_value 需要返回 -1
type MaxQueue struct {
	queue *list.List
	deque *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: list.New(),
		deque: list.New(), //维护最大值
	}
}

func (this *MaxQueue) Max_value() int {
	if this.deque.Len() <= 0 {
		return -1
	}
	return this.deque.Front().Value.(int)
}

func (this *MaxQueue) Push_back(value int) {

	this.queue.PushBack(value)
	//【循环】清除掉小于value的值
	for this.deque.Len() > 0 && this.deque.Back().Value.(int) < value {
		this.deque.Remove(this.deque.Back())
	}
	this.deque.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.queue.Len() <= 0 {
		return -1
	}
	//保持两个队列同步
	if this.deque.Front().Value == this.queue.Front().Value {
		this.deque.Remove(this.deque.Front())
	}
	return this.queue.Remove(this.queue.Front()).(int)
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
