package main

func BTreeMin(root *TreeNode) *TreeNode {
	// in a BST the last node at the left branch is alway the lowest node
	// if root == nil {
	// 	return  nil
	// }
	
	// current := root
	// for current.Left != nil {
	// 	current = current.Left
	// }
	// return current

	// another style 
	if root == nil {
		return nil
	}

	lowest := root
	leftMin := BTreeMin(root.Left)
	if leftMin != nil && leftMin.Data < lowest.Data {
		lowest = leftMin
	}

	rightmin := BTreeMin(root.Right)
	if rightmin != nil && rightmin.Data < lowest.Data {
		lowest = rightmin
	}
	return lowest
}