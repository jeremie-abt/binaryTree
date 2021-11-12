package second_tree

import "reflect"

type node struct {
	parent *node // -> 4
	left   *node // -> 3
	right  *node // -> 5

	value interface{}
}

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()                       }

/*
Private interface
 */

func newNode(val interface{}, parent *node) (*node, error) {

	// TODO
	return &node{
		value: val,
		parent: parent,
	}, nil
}

func getRootNode() (*node, error) {
	return newNode(nil, nil)
}

func (tr *Tree) insert(val interface{}) error {
	var err error = nil
	nd := tr.rootNode

	if nd == nil {
		tr.rootNode, err = getRootNode()
		tr.rootNode.value = val
		return err
	}

	curNode := tr.rootNode
	for {
		if tr.shouldGoLeft(curNode, val) {
			if curNode.left == nil {
				curNode.left, err = newNode(val, curNode)
				break
			}
			curNode = curNode.left
		} else {
			if curNode.right == nil {
				curNode.right, err = newNode(val, curNode)
				break
			}
			curNode = curNode.right
		}
	}
	return err
}

func (tr *Tree) shouldGoLeft(nd *node, val interface{}) bool {

	GreaterThan := tr.GTFunc
	if GreaterThan(nd.value, val) {
		return true
	}
	return false
}
	/*
	   Public Interface
 */

type Tree struct {
	rootNode *node

	GTFunc func(interface{}, interface{}) bool
}

func NewTree(GTFunc func(interface{}, interface{}) bool, val ...interface{}) (*Tree, error) {

	var err error = nil

	if len(val) == 0 {
		return &Tree{
			rootNode: nil,
			GTFunc:   GTFunc,
		}, err
	}
	if len(val) != 1 {
		panic("not implemented yet")
	}
	node, err := newNode(val[0], nil)
	return &Tree{
		rootNode:       node,
		GTFunc: GTFunc,
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
