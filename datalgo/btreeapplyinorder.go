package main

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// base case to stop recursiveness
	if root == nil {
		return
	}
	
	// go depth through the left
	BTreeApplyInorder(root.Left, f)

	// print data if base is reached
	f(root.Data)

	// go depth through the right
	BTreeApplyInorder(root.Right, f)
}