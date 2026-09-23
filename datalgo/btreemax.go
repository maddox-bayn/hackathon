package main

func BTreeMax(root *TreeNode) *TreeNode {
	// // base case
	// if root == nil {
	// 	return nil
	// }

	// heighest := root
	// leftMax := BTreeMax(root.Left)
	// if leftMax != nil && leftMax.Data > heighest.Data {
	// 	heighest = leftMax
	// }

	// // check right node branch
	// rightMax := BTreeMax(root.Right)
	// if rightMax != nil && rightMax.Data > heighest.Data {
	// 	heighest = rightMax
	// }
	// return heighest

	// another simple solution
	// in Bst the largest value is always the the rightmost node
	if root == nil {
		return nil
	}
	current := root
	for current.Right != nil {
		current = current.Right

	}
	return current
}
