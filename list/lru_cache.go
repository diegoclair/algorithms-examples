package list

import (
	"container/list"
	"fmt"
)

func LruCache() {
	resp := NewLruCache(3)
	resp.Put(1, 1)
	fmt.Println("Add 1,1: ")
	resp.Print()
	resp.Put(2, 2)
	fmt.Println("Add 2,2: ")
	resp.Print()
	resp.Put(1, 1)
	fmt.Println("Add 1,1: ")
	resp.Print()
	resp.Put(2, 2)
	fmt.Println("Add 2,2: ")
	resp.Print()
	resp.Put(3, 3)
	fmt.Println("Add 3,3: ")
	resp.Print()
	resp.Put(4, 4)
	fmt.Println("Add 4,4: ")
	resp.Print()
	res := resp.Get(4)
	fmt.Println("Get 4: ", res)
	resp.Print()
	res = resp.Get(3)
	fmt.Println("Get 3: ", res)
	resp.Print()
	res = resp.Get(2)
	fmt.Println("Get 2: ", res)
	resp.Print()
	res = resp.Get(1)
	fmt.Println("Get 1: ", res)
	resp.Print()
	resp.Put(5, 5)
	fmt.Println("Add 5,5: ")
	resp.Print()
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	linkList *list.List
}

func NewLruCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element, capacity),
		linkList: list.New(),
	}
}

func (l *LRUCache) Get(key int) int {
	result := -1
	if e, ok := l.cache[key]; ok {
		result = e.Value.([]int)[1]
		l.linkList.MoveToFront(e)
	}

	return result
}

func (l *LRUCache) Put(key int, value int) {
	if e, ok := l.cache[key]; ok {
		l.linkList.Remove(e)
		delete(l.cache, key)
	}

	l.cache[key] = l.linkList.PushFront([]int{key, value})

	if l.linkList.Len() > l.capacity {
		lastElement := l.linkList.Back()
		v := l.linkList.Remove(lastElement)
		delete(l.cache, v.([]int)[0])
	}
}

func (l *LRUCache) Print() {
	print := [][]int{}
	for e := l.linkList.Front(); e != nil; e = e.Next() {
		print = append(print, e.Value.([]int))
	}
	fmt.Println(print)
}
