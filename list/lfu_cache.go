package list

import "container/list"

type Element struct {
	key   int
	value int
	freq  int
}

type LFUCache struct {
	cache    map[int]*list.Element
	freq     map[int]*list.List // freq is a map of frequency to a list of elements, the key is the frequency and the value is a list of elements with that frequency
	capacity int
	minFreq  int
}

func NewLfuCache(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		freq:     make(map[int]*list.List),
		minFreq:  0,
	}
}

func (l *LFUCache) Get(key int) int {
	element, ok := l.cache[key]
	if !ok {
		return -1
	}

	e := element.Value.(*Element)

	freqList := l.freq[e.freq]
	freqList.Remove(element)

	if freqList.Len() == 0 && e.freq == l.minFreq {
		l.minFreq++
	}

	e.freq++
	newList, ok := l.freq[e.freq]
	if !ok {
		newList = list.New()
		l.freq[e.freq] = newList
	}

	l.cache[key] = newList.PushFront(e)

	return e.value
}

func (l *LFUCache) Put(key int, value int) {
	if l.capacity == 0 {
		return
	}

	if e, ok := l.cache[key]; ok {
		element := e.Value.(*Element)
		element.value = value
		l.Get(key)
		return
	}

	if l.capacity == len(l.cache) {
		minList := l.freq[l.minFreq]
		lastUsedElement := minList.Back()
		delete(l.cache, lastUsedElement.Value.(*Element).key)
		minList.Remove(lastUsedElement)
	}

	newList, ok := l.freq[1]
	if !ok {
		newList = list.New()
		l.freq[1] = newList
	}

	l.minFreq = 1
	newElement := &Element{key, value, 1}

	element := newList.PushFront(newElement)
	l.cache[key] = element
}
