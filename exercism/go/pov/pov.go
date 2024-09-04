package pov

type Tree struct {
	V string
	P *Tree
	C []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	tree := &Tree{V: value}
	for _, child := range children {
		child.P = tree
		tree.C = append(tree.C, child)
	}
	return tree
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.V
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.C
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	fromNode := tr.findNode(from)
	if fromNode == nil {
		return nil
	}
	if fromNode.P == nil {
		return fromNode
	}

	// start at the fromNode to reassign references
	var np, n, op *Tree = nil, fromNode, fromNode.P
	for n != nil {
		n.P = np
		if op != nil {
			n.C = append(n.C, op)
		}
		// remove np as a child from n
		if np != nil {
			for i, child := range n.C {
				if child == np {
					n.C = append(n.C[:i], n.C[i+1:]...)
					break
				}
			}
		}

		// move one level up
		n, np = op, n
		if op != nil {
			op = op.P
		} else {
			op = nil
		}
	}
	return fromNode
}

func (tr *Tree) findNode(v string) *Tree {
	if tr.Value() == v {
		return tr
	}
	for _, child := range tr.C {
		if node := child.findNode(v); node != nil {
			return node
		}
	}
	return nil
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	if tr = tr.FromPov(from); tr != nil {
		return tr.findPathTo(to)
	}
	return nil
}

func (tr *Tree) findPathTo(to string) []string {
	if tr.V == to {
		return []string{to}
	}
	for _, child := range tr.C {
		if p := child.findPathTo(to); p != nil {
			return append([]string{tr.V}, p...)
		}
	}
	return nil
}
