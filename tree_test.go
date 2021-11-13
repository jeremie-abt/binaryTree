package second_tree

import (
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func TestNewTree(t *testing.T) {
	tr, _ := NewTree(intComparisonFunc, 8, 9, 12, 3, 4)

	assert.Assert(t, tr.rootNode.right.value.(int) == 9)
	assert.Assert(t, tr.rootNode.right.right.value.(int) == 12)
	assert.Assert(t, tr.rootNode.left.value.(int) == 3)
	assert.Assert(t, tr.rootNode.left.right.value.(int) == 4)

	_, err := NewTree(intComparisonFunc, 8, 9, 12, 3, "bad data type")
	assert.Error(t, err, "binary tree can not contain different data type")
}

func TestGetAsList(t *testing.T) {
	tr := getBasicIntTree()
	assert.DeepEqual(t, tr.GetAsList(), []interface{}{1, 2, 3, 4, 8, 12, 15, 19})

	tr, _ = NewTree(intComparisonFunc)
	assert.Assert(t, cmp.Nil(tr.GetAsList()))
}

func TestInsertNode(t *testing.T) {
	tr := getBasicIntTree()

	assert.Assert(t, tr.rootNode.value.(int) == 8)
	assert.Assert(t, tr.rootNode.right.value.(int) == 12)
	assert.Assert(t, tr.rootNode.left.value.(int) == 2)
	assert.Assert(t, tr.rootNode.left.right.value.(int) == 3)
	assert.Assert(t, tr.rootNode.right.right.value.(int) == 15)
	assert.Assert(t, tr.rootNode.right.right.right.value.(int) == 19)
	assert.Assert(t, tr.rootNode.left.left.value.(int) == 1)

	err := tr.Insert("bad data type")
	assert.Error(t, err, "binary tree can not contain different data type")
}

func TestInsertSpecificFunc(t *testing.T) {
	tr := getBasicStrTree()

	assert.Assert(t, tr.rootNode.value.(string) == "bonjour")
	assert.Assert(t, tr.rootNode.right.value.(string) == "bonjour a tous")
	assert.Assert(t, tr.rootNode.left.value.(string) == "bo")
	assert.Assert(t, tr.rootNode.left.right.value.(string) == "bon")
	assert.Assert(t, tr.rootNode.right.right.value.(string) == "salut a tous les amis")
	assert.Assert(t, tr.rootNode.right.right.right.value.(string) == "salut salut a tous les ami")
	assert.Assert(t, tr.rootNode.left.left.value.(string) == "s")
}

func TestDeleteLeaf(t *testing.T) {
	tr := getBasicIntTree()

	_ = tr.delete(1)
	_ = tr.delete(4)
	_ = tr.delete(19)
	assert.Assert(t, cmp.Nil(tr.rootNode.left.left))
	assert.Assert(t, cmp.Nil(tr.rootNode.left.right.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.right.right.right))

	_ = tr.delete(15)
	_ = tr.delete(3)
	assert.Assert(t, cmp.Nil(tr.rootNode.right.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.left.right))

	_ = tr.delete(12)
	_ = tr.delete(2)
	assert.Assert(t, cmp.Nil(tr.rootNode.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.left))
}

func TestDeleteRoot(t *testing.T) {
	tr, _ := NewTree(intComparisonFunc)

	_ = tr.insert(8)
	_ = tr.delete(8)

	assert.Assert(t, cmp.Nil(tr.rootNode))

	_ = tr.insert(8)
	_ = tr.insert(8)
	_ = tr.delete(8)
	assert.Assert(t, tr.rootNode.value.(int) == 8)
	_ = tr.delete(8)
	assert.Assert(t, cmp.Nil(tr.rootNode))

	_ = tr.insert(8)
	_ = tr.insert(6)
	_ = tr.insert(15)
	err := tr.delete(78)
	assert.Error(t, err, "value not exist in binary tree")
}

func TestRemoveOneChildNode(t *testing.T) {
	tr := getBasicIntTree()

	_ = tr.delete(12)
	_ = tr.delete(3)

	assert.Assert(t, tr.rootNode.right.value.(int) == 15)
	assert.Assert(t, tr.rootNode.right.right.value.(int) == 19)
	assert.Assert(t, tr.rootNode.left.right.value.(int) == 4)
}

func TestRemoveNodeWithChilds(t *testing.T) {
	tr := getBasicIntTree()

	_ = tr.delete(12)
	_ = tr.delete(2)
	assert.Assert(t, tr.rootNode.right.value.(int) != 12)
	assert.Assert(t, tr.rootNode.left.value.(int) != 2)
}

func intComparisonFunc(first interface{}, sec interface{}) bool {
	if first.(int) < sec.(int) {
		return true
	}
	return false
}

func strComparisonFunc(first interface{}, sec interface{}) bool {
	if len(first.(string)) < len(sec.(string)) {
		return true
	}
	return false
}

func getBasicIntTree() *tree {
	/*
							8
						   / \
		                  2   12
		                 / \    \
		                1   3    15
		                     \     \
		                      4     19
	*/
	tr, _ := NewTree(intComparisonFunc)

	_ = tr.Insert(8)
	_ = tr.Insert(12)
	_ = tr.Insert(2)
	_ = tr.Insert(3)
	_ = tr.Insert(15)
	_ = tr.Insert(19)
	_ = tr.Insert(1)
	_ = tr.Insert(4)

	return tr
}

func getBasicStrTree() *tree {
	/*
			"bonjour"
			   / \
		   "bo"    "bonjour a tous"
		   / \             \
		 "s"  "bon"         "salut a tous les amis
										 \
										  "salut salut a tous les amis"
	*/
	tr, _ := NewTree(strComparisonFunc)

	_ = tr.Insert("bonjour")
	_ = tr.Insert("bonjour a tous")
	_ = tr.Insert("bo")
	_ = tr.Insert("bon")
	_ = tr.Insert("salut a tous les amis")
	_ = tr.Insert("salut salut a tous les ami")
	_ = tr.Insert("s")

	return tr
}
