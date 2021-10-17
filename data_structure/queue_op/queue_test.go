package queueop

import (
	"testing"
)

func TestFlt(t *testing.T) {
	l := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
		},
		Child: &Node{
			Val: 3,
		},
	}
	flatten(l)
}
