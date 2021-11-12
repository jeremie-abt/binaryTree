package second_tree

import (
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
	"testing"
)

func TestInsertNode(t *testing.T) {
	tr, _ := NewTree(intComparisonFunc)
	assert.Assert(t, cmp.Nil(tr.rootNode))

	tr.insert(8)
	tr.insert(12)
	tr.insert(2)
	tr.insert(3)
	tr.insert(15)
	tr.insert(19)
	tr.insert(1)

	assert.Assert(t, tr.rootNode.value.(int) == 8)
	assert.Assert(t, tr.rootNode.right.value.(int) == 12)
	assert.Assert(t, tr.rootNode.left.value.(int) == 2)
	assert.Assert(t, tr.rootNode.left.right.value.(int) == 3)
	assert.Assert(t, tr.rootNode.right.right.value.(int) == 15)
	assert.Assert(t, tr.rootNode.right.right.right.value.(int) == 19)
	assert.Assert(t, tr.rootNode.left.left.value.(int) == 1)
}


func TestInsertSpecificFunc(t *testing.T) {
	tr, _ := NewTree(func(first interface{}, sec interface{}) bool {

		if len(first.(string)) < len(sec.(string)) {
			return true
		}
		return false
	})
	assert.Assert(t, cmp.Nil(tr.rootNode))

	tr.insert("bonjour")
	tr.insert("bonjour a tous")
	tr.insert("bo")
	tr.insert("bon")
	tr.insert("salut a tous les amis")
	tr.insert("salut salut a tous les ami")
	tr.insert("s")

	assert.Assert(t, tr.rootNode.value.(string) == "bonjour")
	assert.Assert(t, tr.rootNode.right.value.(string) == "bonjour a tous")
	assert.Assert(t, tr.rootNode.left.value.(string) == "bo")
	assert.Assert(t, tr.rootNode.left.right.value.(string) == "bon")
	assert.Assert(t, tr.rootNode.right.right.value.(string) == "salut a tous les amis")
	assert.Assert(t, tr.rootNode.right.right.right.value.(string) == "salut salut a tous les ami")
	assert.Assert(t, tr.rootNode.left.left.value.(string) == "s")
}

func TestDeleteLeaf(t *testing.T) {
	/*
						8
					   / \
	                  2   12
	                 / \    \
	                1   3    15
	                           \
	                            19
	 */
	tr, _ := NewTree(intComparisonFunc)
	assert.Assert(t, cmp.Nil(tr.rootNode))

	tr.insert(8)
	tr.insert(12)
	tr.insert(2)
	tr.insert(3)
	tr.insert(15)
	tr.insert(19)
	tr.insert(1)

	tr.delete(1)
	tr.delete(3)
	tr.delete(19)
	assert.Assert(t, cmp.Nil(tr.rootNode.left.left))
	assert.Assert(t, cmp.Nil(tr.rootNode.left.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.right.right.right))

	tr.delete(15)
	tr.delete(2)
	assert.Assert(t, cmp.Nil(tr.rootNode.right.right))
	assert.Assert(t, cmp.Nil(tr.rootNode.left))

	tr.delete(12)
	assert.Assert(t, cmp.Nil(tr.rootNode.right))
}

func TestDeleteRoot(t *testing.T){
	tr, _ := NewTree(func(first interface{}, sec interface{}) bool {

		if first.(int) < sec.(int) {
			return true
		}
		return false
	})

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
	tr, _ := NewTree(intComparisonFunc)

	tr.insert(8)
	tr.insert(12)
	tr.insert(2)
	tr.insert(3)
	tr.insert(15)
	tr.insert(19)
	tr.insert(1)
	tr.insert(4)
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
	tr, _ := NewTree(intComparisonFunc)

	tr.insert(8)
	tr.insert(12)
	tr.insert(9)
	tr.insert(15)
	tr.insert(11)
	tr.insert(19)
	tr.insert(2)
	tr.insert(1)
	tr.insert(3)
	tr.insert(4)

	tr.delete(12)
	tr.delete(2)
	assert.Assert(t, tr.rootNode.right.value.(int) != 12)
	assert.Assert(t, tr.rootNode.left.value.(int) != 2)
}

//func TestNewTreeInterface(t *testing.T) {
//	firstTree, err1 := NewNode("test", "test2", "test3")
//	secTree, err2 := NewNode([]string{"test", "test2", "test3"})
//	thirdTree, err3 := NewNode("test", "test2")
//	fourthTree, err4 := NewNode(6, 13, 15, 4, 2)
//	fifthTree, err5 := NewNode(6, 15, 13, 2, 4)
//
//	assert.DeepEqual(t, firstTree, secTree)
//	assert.Check(t, !reflect.DeepEqual(firstTree, thirdTree))
//	assert.DeepEqual(t, fourthTree, fifthTree)
//	assert.NilError(t, err1)
//	assert.NilError(t, err2)
//	assert.NilError(t, err3)
//	assert.NilError(t, err4)
//	assert.NilError(t, err5)
//}

// linked to TestEnforceTyping function.
//type firstCustomType struct {
//	value int
//}

// linked to TestEnforceTyping function.
//type secondCustomType struct {
//	value int
//}

// linked to TestEnforceTyping function.
//type interfaceTest interface{}

//func TestEnforceTyping(t *testing.T) {
	// Values for a binary tree can be of different types.
	// This test verify that statement hold true.

	//_, err := NewNode("test", "test2", 5)
	//assert.Error(t, err, "tree can't hold different type of data")

	//_, err = NewNode(interfaceTest(firstCustomType{value: 5}), interfaceTest(secondCustomType{value: 5}))
	//assert.Check(t, err, "tree can't hold different type of data")

	// tree must accept interface values with same concrete type
	//_, err = NewNode(interfaceTest(firstCustomType{value: 6}), interfaceTest(firstCustomType{value: 5}))
	//assert.NilError(t, err)

	//testedTree, _ := NewNode()
	//err = testedTree.Insert(5, 4)
	//assert.NilError(t, err)
	//err = testedTree.Insert("bad string data type")
	//assert.Error(t, err, "tree can't hold different type of data")
//}

//func TestInsert(t *testing.T) {
//
//	treeConstructViaInsert, _ := NewTree()
//	err := treeConstructViaInsert.Insert(8, 9)
//	assert.NilError(t, err)
//	err = treeConstructViaInsert.Insert(15, 2)
//	assert.NilError(t, err)
//	err = treeConstructViaInsert.Insert(2, 25)
//	assert.NilError(t, err)
//
//	treeConstructViaConstructor, err := NewTree(8, 9, 15, 2, 2, 25)
//	assert.NilError(t, err)
//	assert.DeepEqual(t, treeConstructViaConstructor, treeConstructViaInsert)
//
//	err = treeConstructViaInsert.Insert("bad date type")
//	assert.Error(t, err, "tree can't hold different type of data")
//}
//
//func TestDelete(t *testing.T) {
//
//	testedTree, err := NewTree(8, 9, 15, 2, 2, 25)
//	treeWithValueDeleted, err1 := NewTree(8, 9, 2, 25)
//	assert.NilError(t, err)
//	assert.NilError(t, err1)
//
//	err = testedTree.Delete(15, 2)
//	assert.NilError(t, err)
//	assert.DeepEqual(t, testedTree, treeWithValueDeleted)
//
//	err = testedTree.Delete(1568)
//	assert.Error(t, err, "value doest not exist in this tree")
//	err = testedTree.Delete("bad data type")
//	assert.Error(t, err, "tree can't hold different type of data")
//}

//func TestGetAsList(t *testing.T) {
//	testedTree, _ := NewTree(8, 5, 6, 9, 78, 45, 6)
//
//	assert.DeepEqual(t, []int{5, 6, 6, 8, 9, 45, 78}, testedTree.GetAsList())
//	testedTree.AddTriFunc(func(first interface{}, sec interface{}) int {
//		this func should return 1 if the first element is inferior to the second,
//		0 otherwise
		//if first.(int) > sec.(int) {
		//	I want to sort it decreasingly, so the cond is reversed.
			//return 1
		//}
		//return 0
	//})
	//assert.DeepEqual(t, []int{78, 45, 9, 8, 6, 6, 5}, testedTree.GetAsList())
//}

func intComparisonFunc(first interface{}, sec interface{}) bool {
	if first.(int) < sec.(int) {
		return true
	}
	return false
}

