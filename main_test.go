package main

import (
	"testing"

	"github.com/zeebo/assert"
	"golang.org/x/tour/tree"
)

func TestMain(t *testing.T) {
	main()

	tests := []struct {
		tree1    *tree.Tree
		tree2    *tree.Tree
		expected bool
	}{
		{
			tree1:    tree.New(1),
			tree2:    tree.New(1),
			expected: true,
		},
		{
			tree1:    tree.New(1),
			tree2:    tree.New(2),
			expected: false,
		},
		{
			tree1:    tree.New(2),
			tree2:    tree.New(1),
			expected: false,
		},
	}

	for _, test := range tests {
		got := Same(test.tree1, test.tree2)
		assert.Equal(t, got, test.expected)
	}
}
