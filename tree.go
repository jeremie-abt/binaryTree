package second_tree

import (
	"errors"
	"reflect"
)

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

func (nd *node) isLeaf () bool {
	if nd.left == nil && nd.right == nil {
		return true
	}
	return false
}

func (nd *node) hasOneChild() bool {
	if (nd.left != nil && nd.right == nil) || (nd.right != nil && nd.left == nil) {
		return true
	}
	return false
}

func (nd *node) supprNode() error {
	var err error = nil

	if nd.isLeaf() {
		// cut the leaf
		if nd.parent.left == nd {
			nd.parent.left = nil
		} else {
			nd.parent.right = nil
		}
	} else if nd.hasOneChild() {
		// copier le fils
		var nodeToCopy *node
		if nd.left != nil {
			nodeToCopy = nd.left
		} else {
			nodeToCopy = nd.right
		}
		nd.left = nodeToCopy.left
		nd.right = nodeToCopy.right
		nd.value = nodeToCopy.value
		return err
	} else {
		// suppr node with two childs
		var curNode *node = nd.right
		for {
			if curNode.left == nil {
				break
			}
			curNode = curNode.left
		}
		// curNode est le node a replace
		nd.value = curNode.value
		nd.right = curNode.right
	}
	return err
}

func (tr *Tree) delete(val interface {}) error {
	var err error = nil
	var supprRootNode bool = false

	curNode := tr.rootNode
	for {
		if tr.shouldGoLeft(curNode, val) {
			if curNode.left == nil {
				// il faut le faire avant curNode = curNode.left sinon c'est pas bon
				supprRootNode = true
				break
			}
			curNode = curNode.left
			if reflect.DeepEqual(val, curNode.value) {
				return curNode.supprNode()
			}
		} else {
			if curNode.right == nil {
				supprRootNode = true
				break
			}
			curNode = curNode.right
			if reflect.DeepEqual(val, curNode.value) {
				return curNode.supprNode()
			}
		}
	}
	if supprRootNode == true && reflect.DeepEqual(val, tr.rootNode.value) {
		tr.rootNode = nil
	}
	if supprRootNode == true {
		err = errors.New("value not exist in binary tree")
	}
	return err
}

func (tr *Tree) shouldGoLeft(nd *node, val interface{}) bool {

	LessThanFunc := tr.LTFunc
	if LessThanFunc(val, nd.value) {
		return true
	}
	return false
}
	/*
	   Public Interface
 */

type Tree struct {
	rootNode *node

	LTFunc func(interface{}, interface{}) bool
}

func NewTree(LTFunc func(interface{}, interface{}) bool, val ...interface{}) (*Tree, error) {

	var err error = nil

	if len(val) == 0 {
		return &Tree{
			rootNode: nil,
			LTFunc:   LTFunc,
		}, err
	}
	if len(val) != 1 {
		panic("not implemented yet")
	}
	node, err := newNode(val[0], nil)
	return &Tree{
		rootNode: node,
		LTFunc:   LTFunc,
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
