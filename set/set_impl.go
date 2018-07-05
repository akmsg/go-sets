package set

import "sync"

type setImpl struct {
	m map[interface{}]struct{} // struct{} doesn't take up space
	l sync.RWMutex             // use this to make sure concurrent writing is find, use only when
}

// New creates a new set and populate it with items
func New(items ...interface{}) *Set {
	s := &setImpl{}

	for _, item := range items {
		s.m[item] = dummyValue
	}
	return s
}

// Add adds new items in the set
func (s *setImpl) Add(items ...interface{}) bool {
	if 0 >= len(items) {
		return false
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, item := range items {
		s.m[item] = dummyValue
	}
	return true
}

// Remove removes items from the Set
func (s *setImpl) Remove(items ...interface{}) bool {
	if 0 >= len(items) {
		return false
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, item := range items {
		delete(s.m, item)
	}
	return true
}
