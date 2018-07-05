package set

import (
	"sync"
)

type Set struct {
	m map[interface{}]struct{} // struct{} doesn't take up space
	l sync.RWMutex             // use this to make sure concurrent writing is find, use only when
}

// Add adds new items in the set
func (s *Set) Add(elements ...interface{}) {
	if 0 >= len(elements) {
		return
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, element := range elements {
		s.m[element] = dummyValue
	}
}

// Remove removes items from the Set
func (s *Set) Remove(elements ...interface{}) {
	if 0 >= len(elements) {
		return
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, element := range elements {
		delete(s.m, element)
	}
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
func (s *Set) Empty() {
	if 0 >= len(s.m) {
		return
	}

	s.l.Lock()
	defer s.l.Unlock()

	s.m = make(map[interface{}]struct{})
}

// Contains will return true if the map has the key
func (s *Set) Contains(element interface{}) bool {
	_, status := s.m[element]
	return status
}
