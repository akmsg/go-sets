package go_sets

import (
	"testing"

	"github.com/anandmishra01/go-sets/set"
)

func TestCreation(t *testing.T) {
	s := set.New("1", 2, 3, 4, 5)

	t.Log(s.List())

	s.Add(1, 5, s)

	t.Log(s.List())

	s.Remove(2)

	t.Log(s.List())
}
