package go_sets

import (
	"testing"

	"github.com/anandmishra01/go-sets/set"
	"github.com/stretchr/testify/require"
)

func TestCreation(t *testing.T) {

	s := set.New("1", 2, 3, 4, 5)

	require.True(t, s.Contains("1"))
	require.True(t, s.Contains(4))

	s.Add(1, 56)
	require.True(t, s.Contains(56))

	s.Remove(1)
	require.False(t, s.Contains(1))

}
