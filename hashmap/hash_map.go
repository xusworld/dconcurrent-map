package hashmap

import (
	"sync"
)

type HashMap interface {
	Set(key string, val interface{})

	Get(key string) interface{}

	Del(key string)

	Len() int

	ForEach(func(key string, val interface{}) bool)
}

// ConcurrentMap
type ConcurrentMap struct {
	segments   []*Segment
	shardCount int
}

func NewConcurrentMap(segmentNums int) *ConcurrentMap {
	cm := &ConcurrentMap{}
	cm.shardCount = segmentNums
	for i := 0; i < segmentNums; i++ {
		segment := newSegment()
		cm.segments = append(cm.segments, segment)
	}
	return cm
}

func (cm *ConcurrentMap) GetSegment(key string) *Segment {
	return cm.segments[uint(fnv32(key))%uint(32)]
}

func (cm *ConcurrentMap) Set(key string, val interface{}) {
	segment := cm.GetSegment(key)
	segment.Set(key, val)
}

func (cm *ConcurrentMap) Get(key string) interface{} {
	segment := cm.GetSegment(key)

	segment.RLock()
	val := segment.items[key]
	segment.RUnlock()
	return val
}

func (cm *ConcurrentMap) Del(key string) {
	segment := cm.GetSegment(key)
	segment.Del(key)
}

func (cm *ConcurrentMap) Len() int {
	size := 0
	for _, segment := range cm.segments {
		size += segment.Size()
	}
	return size
}

func (cm *ConcurrentMap) ForEach(func(key string, val interface{}) bool) {
	// TODO
}

func (cm *ConcurrentMap) SetShardCount(count int) {
	cm.shardCount = count
}

func (cm *ConcurrentMap) ShardCount() int {
	return cm.shardCount
}

type Segment struct {
	items map[string]interface{}
	size  int
	sync.RWMutex
}

func newSegment() *Segment {
	s := &Segment{}
	s.items = make(map[string]interface{})
	s.size = 0

	return s
}

func (s *Segment) Set(key string, val interface{}) {
	s.Lock()
	s.items[key] = val
	s.size++
	s.Unlock()
}

func (s *Segment) Del(key string) {
	s.Lock()
	delete(s.items, key)
	s.size--
	s.Unlock()
}

func (s *Segment) Size() int {
	s.RLock()
	size := s.size
	s.Unlock()
	return size
}

func (s *Segment) Clear() {
	s.Lock()
	s.items = make(map[string]interface{})
	s.size = 0
	s.Unlock()
}

func (s *Segment) ForEach(_ func(key string, val interface{}) bool) {
	panic("not implement yet!")
}
