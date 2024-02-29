package hot

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*LinkedNode
	head, tail *LinkedNode
}

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}

func InitLinkedNode(key, value int) *LinkedNode {
	return &LinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*LinkedNode),
		head:     InitLinkedNode(0, 0),
		tail:     InitLinkedNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		// 移动到头部
		this.moveToHead(v)
		// 返回值
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		// 新建节点
		node := InitLinkedNode(key, value)
		// 长度增加
		this.size++
		// 放入缓存
		this.cache[key] = node
		// 移动到头部
		this.addToHead(node)
		// 判断如果长度大于容量
		if this.size > this.capacity {
			// 移除尾部节点
			rmNode := this.removeTail()
			// 长度-1
			this.size--
			// 删除缓存
			delete(this.cache, rmNode.key)
		}
	} else {
		v.value = value
		this.moveToHead(v)
	}
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *LinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *LinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node

}

func (this *LRUCache) removeTail() *LinkedNode {
	rmNode := this.tail.prev
	this.removeNode(rmNode)
	return rmNode
}
