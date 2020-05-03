package tree

import (
	"fmt"
	"sort"
)

// Record is in the input case.
type Record struct {
	ID, Parent int
}

// Node is the classification style.
type Node struct {
	ID       int
	Children []*Node
}

// Build takes in Record, sorts it, checks for errors, appends records to nodes and returns Node.
func Build(r []Record) (*Node, error) {
	if len(r) == 0 {
		return nil, nil
	}

	tree := make([]*Node, len(r))

	sort.Slice(r, func(i, j int) bool {
		return r[i].ID < r[j].ID
	})

	for i, rec := range r {
		if rec.ID >= len(r) {
			return nil, fmt.Errorf("out of range ID to index: %d", rec.ID)
		}
		if rec.ID != i || rec.Parent > rec.ID || rec.ID > 0 && rec.Parent == rec.ID {
			return nil, fmt.Errorf("not in sequence or has bad parent: %#v -> %#v", rec.ID, rec.Parent)
		}

		node := &Node{rec.ID, nil} // create node
		tree[rec.ID] = node        // append it to index

		if rec.ID != rec.Parent {
			parent := tree[rec.Parent]                      // find Parent
			parent.Children = append(parent.Children, node) // append to Parent
		}
	}

	return tree[0], nil
}