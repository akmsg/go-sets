package set

import (
	"sync"
)

type Set struct {
	m map[interface{}]struct{} // struct{} doesn't take up space
	l sync.RWMutex             // use this to make sure concurrent writing is find, use only when
}

// Add adds new items in the set
func (s *Set) Add(elements ...interface{}) bool {
	if 0 >= len(elements) {
		return false
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, element := range elements {
		s.m[element] = dummyValue
	}
	return true
}

// Remove removes items from the Set
func (s *Set) Remove(elements ...interface{}) bool {
	if 0 >= len(elements) {
		return false
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, element := range elements {
		delete(s.m, element)
	}
	return true
}

// List returns an enumeration of set elements
func (s *Set) List() []interface{} {
	if 0 >= len(s.m) {
		return nil
	}
	var elements []interface{}
	for element, _ := range s.m {
		elements = append(elements, element)
	}
	return elements
}

// Empty removes all the elements from the list
func (s *Set) Empty() bool {
	if 0 >= len(s.m) {
		return false
	}

	s.l.Lock()
	defer s.l.Unlock()

	s.m = make(map[interface{}]struct{})
	return true
}
