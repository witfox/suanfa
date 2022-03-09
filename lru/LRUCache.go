package lru

/**
* 最近最少使用缓存 LRU
 */

type LRUCache struct {
	size               int
	capacity           int
	cache              map[int]*DLinkedNode
	headNode, tailNode *DLinkedNode
}

//双向链表
type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*DLinkedNode),
		capacity: capacity,
		headNode: initDLinkedNode(0, 0),
		tailNode: initDLinkedNode(0, 0),
	}
	l.headNode.next = l.tailNode
	l.tailNode.pre = l.headNode
	return l
}

func (lc *LRUCache) Get(key int) int {
	if _, ok := lc.cache[key]; !ok {
		return -1
	}

	//将key添加到链表前面
	lc.moveToHead(lc.cache[key])

	return lc.cache[key].value
}

func (lc *LRUCache) put(key, value int) {
	if _, ok := lc.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		lc.cache[key] = node
		lc.addToHead(node)
		lc.size++
		//是否超出容量
		if lc.size > lc.capacity {
			removed := lc.removeTail()
			delete(lc.cache, removed.key)
			lc.size--
		}
	} else {
		//如果已经存在,则更新
		node := lc.cache[key]
		node.value = value
		lc.moveToHead(node)
	}
}

func (lc *LRUCache) addToHead(node *DLinkedNode) {
	node.next = lc.headNode.next
	node.pre = lc.headNode

	lc.headNode.next.pre = node
	lc.headNode.next = node
}

func (lc *LRUCache) removeNode(node *DLinkedNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
}

func (lc *LRUCache) moveToHead(node *DLinkedNode) {
	lc.removeNode(node)
	lc.addToHead(node)
}

func (lc *LRUCache) removeTail() *DLinkedNode {
	node := lc.tailNode.pre
	lc.removeNode(node)
	return node
}
