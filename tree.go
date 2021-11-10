package second_tree

import "errors"

type node struct {
	parent *node // -> 4
	left   *node // -> 3
	right  *node // -> 5

	value interface{}
}

/*
Private interface
 */

func newNode(val interface{}) (*node, error) {

	// TODO
	return &node{
		value: val,
	}, nil
}

func (node *node) addLeftChild (val interface{}) error {
	var err error = nil
	if node.left != nil {
		return errors.New("this node has already left child")
	}
	node.left, err = NewNode(val)
	return err
}

func (node *node) addRightChild (val interface{}) (error) {
	node.right.value = val
	return nil
}

/*
Public Interface
 */

type Tree struct {
	rootNode *node
}

func NewTree(val ...interface{}) (*Tree, error) {

	var err error = nil

	if len(val) != 1 {
		panic("not implemented yet")
	}
	node, err := newNode(val[0])
	return &Tree{
		rootNode: node,
	}, err
}

func (node *Tree) Insert(vals ...interface{}) error {
	return nil
}

func (node *Tree) Delete(vals ...interface{}) error {
	return nil
}

func (node *Tree) GetAsList() []interface{} {
	return nil
}

func (node *Tree) AddTriFunc(func(interface{}, interface{}) int) {}
