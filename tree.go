package second_tree

import (
	"errors"
	"reflect"
)

type Tree struct {
	rootNode *node

	LTFunc func(interface{}, interface{}) bool
}

type node struct {
	parent *node
	left   *node
	right  *node

	value interface{}
}

func NewTree(LTFunc func(interface{}, interface{}) bool, val ...interface{}) (*Tree, error) {

	// TODO: COde and refactor
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

func newNode(val interface{}, parent *node) (*node, error) {

	return &node{
		value:  val,
		parent: parent,
	}, nil
}

func getRootNode() (*node, error) {
	return newNode(nil, nil)
}

func (tr *Tree) Insert(vals ...interface{}) error {
	return nil
}

func (tr *Tree) Delete(vals ...interface{}) error {
	return nil
}

func (tr *Tree) GetAsList() []interface{} {
	return nil
}

func (tr *Tree) AddTriFunc(func(interface{}, interface{}) int) {}

func (tr *Tree) insert(val interface{}) error {
	// TODO: relire ce code pour voir s'il reste clair
	var err error = nil
	var ChildToLookFor **node
	nd := tr.rootNode

	if nd == nil {
		return tr.InstanciateRootNode(val)
	}

	currentNode := tr.rootNode
	for {
		if tr.shouldGoLeft(currentNode, val) {
			ChildToLookFor = &currentNode.left
		} else {
			ChildToLookFor = &currentNode.right
		}
		if *ChildToLookFor == nil {
			*ChildToLookFor, err = newNode(val, currentNode)
			break
		}
		currentNode = *ChildToLookFor
	}
	return err
}

func (tr *Tree) delete(val interface{}) error {
	var childToLookFor *node

	curNode := tr.rootNode
	for {
		if tr.shouldGoLeft(curNode, val) {
			childToLookFor = curNode.left
		} else {
			childToLookFor = curNode.right
		}
		if childToLookFor == nil {
			return tr.tryToDeleteRootNode(val)
		}
		curNode = childToLookFor
		if reflect.DeepEqual(val, curNode.value) {
			return curNode.supprNode()
		}
	}
}

func (tr *Tree) shouldGoLeft(nd *node, val interface{}) bool {

	LessThanFunc := tr.LTFunc
	if LessThanFunc(val, nd.value) {
		return true
	}
	return false
}

func (tr *Tree) tryToDeleteRootNode(val interface{}) error {
	var err error = nil

	if reflect.DeepEqual(val, tr.rootNode.value) {
		tr.rootNode = nil
	} else {
		err = errors.New("value not exist in binary tree")
	}
	return err
}

func (nd *node) supprNode() error {
	var err error = nil

	if nd.isLeaf() {
		return nd.cutLeaf()
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
		var curNode = nd.right
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

func (nd *node) cutLeaf() error {
	if nd.parent.left == nd {
		nd.parent.left = nil
	} else {
		nd.parent.right = nil
	}
	return nil
}

func (tr *Tree) InstanciateRootNode(val interface{}) error {
	var err error = nil
	tr.rootNode, err = getRootNode()
	tr.rootNode.value = val
	return err
}

func (nd *node) isLeaf() bool {
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
