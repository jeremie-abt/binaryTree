package second_tree

import (
	"gotest.tools/assert/cmp"
	"testing"
)
import "reflect"

import "gotest.tools/assert"

/*
Private test for tree

func addLeftChild()  // work only on leaf who has the place to have left child
func addRightChild() // work only on leaf who has the place to have left child

func removeNode() // work on the specific node
*/

func TestAddLeftChild(t *testing.T) {
	node, _ := NewNode(15)
	err := node.addLeftChild(9)
	assert.NilError(t, err)
	err = node.left.addLeftChild(10)
	assert.NilError(t, err)
	err = node.addLeftChild(11)
	assert.Error(t, err, "this node has already left child")

	assert.Assert(t, cmp.Nil(node.right))
	assert.Assert(t, node.value.(int) == 15)
	//assert.Assert(t, node.left.value.(int) == 9)
	//assert.Assert(t, node.left.left.value.(int) == 10)
}

func TestAddRightChild(t *testing.T) {
	node, _ := NewNode(15)
	err := node.addRightChild(9)
	assert.NilError(t, err)
	err = node.right.addRightChild(10)
	assert.NilError(t, err)
	err = node.addRightChild(11)
	assert.Error(t, err, "this node has already right child")

	assert.Assert(t, cmp.Nil(node.right))
	assert.Assert(t, node.value.(int) == 15)
	assert.Assert(t, node.right.value.(int) == 9)
	assert.Assert(t, node.right.right.value.(int) == 10)
}

/*
Test for public interface
*/

func TestNewTreeInterface(t *testing.T) {
	firstTree, err1 := NewNode("test", "test2", "test3")
	secTree, err2 := NewNode([]string{"test", "test2", "test3"})
	thirdTree, err3 := NewNode("test", "test2")
	fourthTree, err4 := NewNode(6, 13, 15, 4, 2)
	fifthTree, err5 := NewNode(6, 15, 13, 2, 4)

	assert.DeepEqual(t, firstTree, secTree)
	assert.Check(t, !reflect.DeepEqual(firstTree, thirdTree))
	assert.DeepEqual(t, fourthTree, fifthTree)
	assert.NilError(t, err1)
	assert.NilError(t, err2)
	assert.NilError(t, err3)
	assert.NilError(t, err4)
	assert.NilError(t, err5)
}

// linked to TestEnforceTyping function.
type firstCustomType struct {
	value int
}

// linked to TestEnforceTyping function.
type secondCustomType struct {
	value int
}

// linked to TestEnforceTyping function.
type interfaceTest interface{}

func TestEnforceTyping(t *testing.T) {
	// Values for a binary tree can be of different types.
	// This test verify that statement hold true.

	_, err := NewNode("test", "test2", 5)
	assert.Error(t, err, "tree can't hold different type of data")

	_, err = NewNode(interfaceTest(firstCustomType{value: 5}), interfaceTest(secondCustomType{value: 5}))
	assert.Check(t, err, "tree can't hold different type of data")

	// tree must accept interface values with same concrete type
	_, err = NewNode(interfaceTest(firstCustomType{value: 6}), interfaceTest(firstCustomType{value: 5}))
	assert.NilError(t, err)

	testedTree, _ := NewNode()
	err = testedTree.Insert(5, 4)
	assert.NilError(t, err)
	err = testedTree.Insert("bad string data type")
	assert.Error(t, err, "tree can't hold different type of data")
}

func TestInsert(t *testing.T) {

	treeConstructViaInsert, _ := NewTree()
	err := treeConstructViaInsert.Insert(8, 9)
	assert.NilError(t, err)
	err = treeConstructViaInsert.Insert(15, 2)
	assert.NilError(t, err)
	err = treeConstructViaInsert.Insert(2, 25)
	assert.NilError(t, err)

	treeConstructViaConstructor, err := NewTree(8, 9, 15, 2, 2, 25)
	assert.NilError(t, err)
	assert.DeepEqual(t, treeConstructViaConstructor, treeConstructViaInsert)

	err = treeConstructViaInsert.Insert("bad date type")
	assert.Error(t, err, "tree can't hold different type of data")
}

func TestDelete(t *testing.T) {

	testedTree, err := NewTree(8, 9, 15, 2, 2, 25)
	treeWithValueDeleted, err1 := NewTree(8, 9, 2, 25)
	assert.NilError(t, err)
	assert.NilError(t, err1)

	err = testedTree.Delete(15, 2)
	assert.NilError(t, err)
	assert.DeepEqual(t, testedTree, treeWithValueDeleted)

	err = testedTree.Delete(1568)
	assert.Error(t, err, "value doest not exist in this tree")
	err = testedTree.Delete("bad data type")
	assert.Error(t, err, "tree can't hold different type of data")
}

func TestGetAsList(t *testing.T) {
	testedTree, _ := NewTree(8, 5, 6, 9, 78, 45, 6)

	assert.DeepEqual(t, []int{5, 6, 6, 8, 9, 45, 78}, testedTree.GetAsList())
	testedTree.AddTriFunc(func(first interface{}, sec interface{}) int {
		// this func should return 1 if the first element is inferior to the second,
		// 0 otherwise
		if first.(int) > sec.(int) {
			// I want to sort it decreasingly, so the cond is reversed.
			return 1
		}
		return 0
	})
	assert.DeepEqual(t, []int{78, 45, 9, 8, 6, 6, 5}, testedTree.GetAsList())
}
