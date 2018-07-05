package set

// New creates a new set and populate it with items
func New(elements ...interface{}) *Set {
	s := &Set{m: make(map[interface{}]struct{})}

	s.Add(elements...)

	return s
}
