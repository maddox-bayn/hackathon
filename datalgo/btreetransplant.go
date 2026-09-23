package main

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	// The standard way to implement a structural pointer update recursively 
	// is to re-assign the left/right pointers during the return phase
	// if the tree is empty nothing o to transplant.
	if root == nil {
		return nil
	}

	// We found the exact node to replace
	if root == node {
		return rplc
	}
	// recurse left and right, relinking the result to maintain the tree structure
	root.Left = BTreeTransplant(root.Left, node, rplc)

	root.Right = BTreeTransplant(root.Right, node, rplc)

	return root
}