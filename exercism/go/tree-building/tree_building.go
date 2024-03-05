package tree

import (
	"errors"
	"sort"
)

var ErrInvalidID = errors.New("invalid ID")
var ErrInvalidParent = errors.New("invalid parent")

// Record is the struct that represents the data in the input.
// ID is the unique identifier of the record.
// Parent is the ID of the parent of the record.
type Record struct {
	ID     int
	Parent int
}

// Node is the struct that represents the data in the tree.
// ID is the unique identifier of the node.
// Children is an array of pointers to the children of the node.
type Node struct {
	ID       int
	Children []*Node
	ParentID int
}

// Build constructs a tree structure from a slice of records.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Create a map to store the nodes
	nodes := make(map[int]*Node)

	// Iterate through the records and create nodes.
	for _, record := range records {
		if record.ID < 0 || record.ID >= len(records) {
			return nil, ErrInvalidID
		}
		if record.ID == 0 && record.Parent != 0 {
			return nil, ErrInvalidParent
		}
		if record.ID > 0 && record.ID <= record.Parent {
			return nil, ErrInvalidParent
		}
		if _, ok := nodes[record.ID]; ok {
			// duplicate node id!
			return nil, ErrInvalidID
		}
		nodes[record.ID] = &Node{ID: record.ID, Children: []*Node{}, ParentID: record.Parent}
	}

	// Iterate through the nodes and link them.
	for id, node := range nodes {
		if id != 0 {
			nodes[node.ParentID].Children = append(nodes[node.ParentID].Children, node)
			sort.Slice(nodes[node.ParentID].Children, func(i, j int) bool {
				return nodes[node.ParentID].Children[i].ID < nodes[node.ParentID].Children[j].ID
			})
		}
	}
	return nodes[0], nil
}
