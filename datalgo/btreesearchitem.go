package main

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	// base case using  recursioin to traverse through the treenode if not found
	if root == nil {
		return nil
	}
	// inorder walk of left child
	leftResult := BTreeSearchItem(root.Left, elem)
	if leftResult != nil {
		return leftResult
	}

	// check through each node
	if root.Data == elem {
		return root
	}
	// inorder walk of right child
	return BTreeSearchItem(root.Right, elem)
}
