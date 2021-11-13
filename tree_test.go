package second_tree

import (
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func TestInsertNode(t *testing.T) {
	tr := getBasicIntTree()

	assert.Assert(t, tr.rootNode.value.(int) == 8)
	assert.Assert(t, tr.rootNode.right.value.(int) == 12)
	assert.Assert(t, tr.rootNode.left.value.(int) == 2)
	assert.Assert(t, tr.rootNode.left.right.value.(int) == 3)
	assert.Assert(t, tr.rootNode.right.right.value.(int) == 15)
	assert.Assert(t, tr.rootNode.right.right.right.value.(int) == 19)
	assert.Assert(t, tr.rootNode.left.left.value.(int) == 1)
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

	tr.delete(1)
	tr.delete(4)
	tr.delete(19)
	assert.Assert(t, cmp.Nil(tr.rootNode.left.left))
	assert.Assert(t, cmp.Nil(tr.rootNode.left.right.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.right.right.right))

	tr.delete(15)
	tr.delete(3)
	assert.Assert(t, cmp.Nil(tr.rootNode.right.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.left.right))

	tr.delete(12)
	tr.delete(2)
	assert.Assert(t, cmp.Nil(tr.rootNode.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.left))
}

func TestDeleteRoot(t *testing.T) {
	tr, _ := NewTree(intComparisonFunc)

	tr.insert(8)
	tr.delete(8)

	assert.Assert(t, cmp.Nil(tr.rootNode))

	tr.insert(8)
	tr.insert(8)
	tr.delete(8)
	assert.Assert(t, tr.rootNode.value.(int) == 8)
	tr.delete(8)
	assert.Assert(t, cmp.Nil(tr.rootNode))

	tr.insert(8)
	tr.insert(6)
	tr.insert(15)
	err := tr.delete(78)
	assert.Error(t, err, "value not exist in binary tree")
}

func TestRemoveOneChildNode(t *testing.T) {
	tr := getBasicIntTree()

	tr.delete(12)
	tr.delete(3)

	assert.Assert(t, tr.rootNode.right.value.(int) == 15)
	assert.Assert(t, tr.rootNode.right.right.value.(int) == 19)
	assert.Assert(t, tr.rootNode.left.right.value.(int) == 4)
}

func TestRemoveNodeWithChilds(t *testing.T) {
	/*
									8
			                   /         \
			                  2           12
			                 / \          / \
		                    1   3        9    15
			                     \        \     \
			                      4        11     19
	*/
	tr := getBasicIntTree()

	tr.delete(12)
	tr.delete(2)
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

func getBasicIntTree() *Tree {
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

	tr.insert(8)
	tr.insert(12)
	tr.insert(2)
	tr.insert(3)
	tr.insert(15)
	tr.insert(19)
	tr.insert(1)
	tr.insert(4)

	return tr
}

func getBasicStrTree() *Tree {
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

	tr.insert("bonjour")
	tr.insert("bonjour a tous")
	tr.insert("bo")
	tr.insert("bon")
	tr.insert("salut a tous les amis")
	tr.insert("salut salut a tous les ami")
	tr.insert("s")

	return tr
}
