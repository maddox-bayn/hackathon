package main

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// handle base case, if root equal nil 
	if root == nil {
		return
	} 

	// print to start from root node
	f(root.Data)

	// walk through to left side node
	BTreeApplyPreorder(root.Left, f)

	// walk through to right side of the node
	BTreeApplyPreorder(root.Right, f)
}