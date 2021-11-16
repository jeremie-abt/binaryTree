tree is a simple package to manipulate binary tree data structure, it permit you to
insert or delete

### Example

```go
// TODO Faire des fonctions de base.
import github.com/jeremie-abt/tree

var bt := tree.NewTree(IncreasingIntCmpFc)
bt.Insert(15)
bt.Insert(18)
bt.Insert(25)
bt.Insert(24)
bt.Insert(2)
bt.Insert(5)
bt.Insert(6)
bt.Insert(8)
bt.Insert(29)
bt.Insert(34)

bt.Delete(18)

myList := bt.GetAsList()
```