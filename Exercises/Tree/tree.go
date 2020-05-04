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

	// Return if array of records is empty
	if len(r) == 0 {
		return nil, nil
	}

	// Create an array to store the nodes.
	//a An array of n-size where n is the length of the array of records
	tree := make([]*Node, len(r))

	// Arrange nodes in the array in order of parentage
	// Sort the array of records with a `Less` function
	// Our less function effectively does a descending order comparison
	sort.Slice(r, func(i, j int) bool {
		return r[i].ID < r[j].ID
	})

	// Looping through the records
	for i, rec := range r {
		// The Id of a record matches the index of the array of records
		// We cannot have an ID >= to length of array of records
		if rec.ID >= len(r) {
			return nil, fmt.Errorf("out of range ID to index: %d", rec.ID)
		}

		// Check each value of the record has the right parent.
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