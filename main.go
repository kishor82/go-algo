package main

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value, data string) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}

		return n.Left.Insert(value, data)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}

	return nil
}

func (n *Node) Find(s string) (string, bool) {
	if n == nil {
		return "", false
	}

	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

// findMax finds the maximum element in a (sub-)tree. Its value replaces the value of the to-be-deleted node.
// Return values: the node itself and its parent node.
func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}

	if n.Right == nil {
		return n, parent
	}

	return n.Right.findMax(n)
}

// replaceNode replaces the parent’s child pointer to n with a pointer to the replacement node.
// parent must not be nil.
func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}

	parent.Right = replacement

	return nil
}

func (n *Node) Delete(s string, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}

	switch {
	case s < n.Value:
		n.Left.Delete(s, n)
	case s > n.Value:
		n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}

		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}

		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
	}

	replacement, replParent := n.Left.findMax(n)

	n.Value = replacement.Value
	n.Data = replParent.Data

	return replParent.Delete(replacement.Value, replParent)
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}

	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

func (t *Tree) Delete(s string) error {
	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}

	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}

	if fakeParent.Right == nil {
		t.Root = nil
	}

	return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func main() {
	// set up a slice of strings
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	// Create tree and fill it from the values.
	tree := &Tree{}

	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value'", values[i], "':", err)
		}
	}

	// Print tree
	fmt.Printf("%+v\n", tree)
	// Print the sorted values
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ":", n.Data, " | ") })
	fmt.Println()

	// Find values
	s := "d"
	fmt.Print("Find node ' ", s, "' : ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	// Delete a value

	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting"+s+": ", err)
	}

	fmt.Print("After deleting '" + s + "':")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ":", n.Data, " | ") })
	fmt.Println()

	// Special case: A single-node tree.

	fmt.Println("single-node tree")
	tree = &Tree{}

	tree.Insert("a", "alpha")
	fmt.Println("After insert:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ":", n.Data, " | ") })
	fmt.Println()

	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ":", n.Data, " | ") })
	fmt.Println()
}
