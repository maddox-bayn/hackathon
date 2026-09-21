package main

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// base case for recursion
	if root == nil {
		return
	}

	// go depth in left node
	BTreeApplyInorder(root.Left, f)

	// fo depth in right node for post order structure before parent node
	BTreeApplyPostorder(root.Right, f)

	// print current node
	f(root.Data)
}