package main

// 哈希表＋双向链表
// 哈希表中key为缓存的key, value为双向链表节点
// 双向链表节点中存储了key-value键值对

type LRUCache struct {
	Cap        int
	Cache      map[int]*ListNode_
	head, tail *ListNode_
}

type ListNode_ struct {
	Key, Value int
	Prev, Next *ListNode_
}

func Constructor(capacity int) LRUCache {
	head := &ListNode_{}
	tail := &ListNode_{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Cap:   capacity,
		Cache: make(map[int]*ListNode_),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if record, ok := this.Cache[key]; ok {
		this.moveToHead(record)
		return record.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if record, ok := this.Cache[key]; ok {
		record.Value = value
		this.moveToHead(record)
		return
	}
	n := len(this.Cache)
	newNode := &ListNode_{Key: key, Value: value}
	this.addToHead(newNode)
	this.Cache[key] = newNode
	if !(this.Cap > n) {
		node := this.deleteTail()
		delete(this.Cache, node.Key)
	}
}

// addToHead 将node添加到缓存链表头部
func (this *LRUCache) addToHead(node *ListNode_) {
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
	node.Prev = this.head
}

// deleteNode 删除node节点
func (this *LRUCache) deleteNode(node *ListNode_) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// moveToHead 将node移动到缓存链表头部
func (this *LRUCache) moveToHead(node *ListNode_) {
	this.deleteNode(node)
	this.addToHead(node)
}

// deleteTail 删除最近最少使用的node
func (this *LRUCache) deleteTail() *ListNode_ {
	node := this.tail.Prev
	this.deleteNode(this.tail.Prev)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
