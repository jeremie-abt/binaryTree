package second_tree

type Tree struct {}

func NewTree(values ...interface{}) (Tree, error) {
	return &tree{}, nil
}



