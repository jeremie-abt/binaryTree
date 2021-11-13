package second_tree

import (
	"errors"
	"reflect"
)

type tree struct {
	rootNode *node

	LTFunc func(interface{}, interface{}) bool
	typeOf reflect.Type
}

type node struct {
	parent *node
	left   *node
	right  *node

	value interface{}
}

func NewTree(LTFunc func(interface{}, interface{}) bool, vals ...interface{}) (*tree, error) {
	var err error = nil

	tr := &tree{
		LTFunc: LTFunc,
	}
	for _, val := range vals {
		err = tr.Insert(val)
		if err != nil {
			break
		}
	}
	return tr, err
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

func (tr *tree) Insert(val interface{}) error {
	if tr.typeOf == nil {
		tr.typeOf = reflect.TypeOf(val)
	}
	if reflect.TypeOf(val) != tr.typeOf {
		return errors.New("binary tree can not contain different data type")
	}
	return tr.insert(val)
}

func (tr *tree) Delete(val interface{}) error {
	if tr.typeOf == nil {
		tr.typeOf = reflect.TypeOf(val)
	}
	if reflect.TypeOf(val) != tr.typeOf {
		return errors.New("binary tree can not contain different data type")
	}
	return tr.delete(val)
}

func generateInorderTraversal(tr *tree) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		var stack = make([]*node, 0)
		var popedVal *node

		cur := tr.rootNode
		for len(stack) > 0 || cur != nil {
			if cur != nil {
				stack = append(stack, cur)
				cur = cur.left
			} else {
				stack, popedVal = stack[:len(stack)-1], stack[len(stack)-1] // Pop
				ch <- popedVal.value
				cur = popedVal.right
			}
		}
	}()
	return ch
}

func (tr *tree) GetAsSlice() []interface{} {
	var rs []interface{}
	consumer := generateInorderTraversal(tr)
	for val := range consumer {
		rs = append(rs, val)
	}
	return rs
}

func (tr *tree) insert(val interface{}) error {
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

func (tr *tree) delete(val interface{}) error {
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

func (tr *tree) shouldGoLeft(nd *node, val interface{}) bool {

	LessThanFunc := tr.LTFunc
	if LessThanFunc(val, nd.value) {
		return true
	}
	return false
}

func (tr *tree) tryToDeleteRootNode(val interface{}) error {
	var err error = nil

	if reflect.DeepEqual(val, tr.rootNode.value) {
		tr.rootNode = nil
	} else {
		err = errors.New("value not exist in binary tree")
	}
	return err
}

func (nd *node) supprNode() error {
	if nd.isLeaf() {
		return nd.cutLeaf()
	} else if nd.hasOneChild() {
		return nd.cutParentOfOneChild()
	} else {
		return nd.cutParentWithChilds()
	}
}

func (nd *node) cutParentWithChilds() error {
	var curNode = nd.right
	for {
		// get next inorder traversal node
		if curNode.left == nil {
			break
		}
		curNode = curNode.left
	}
	nd.value = curNode.value
	nd.right = curNode.right
	return nil
}

func (nd *node) cutParentOfOneChild() error {
	var nodeToCopy *node

	if nd.left != nil {
		nodeToCopy = nd.left
	} else {
		nodeToCopy = nd.right
	}
	nd.left = nodeToCopy.left
	nd.right = nodeToCopy.right
	nd.value = nodeToCopy.value
	return nil
}

func (nd *node) cutLeaf() error {
	if nd.parent.left == nd {
		nd.parent.left = nil
	} else {
		nd.parent.right = nil
	}
	return nil
}

func (tr *tree) InstanciateRootNode(val interface{}) error {
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
