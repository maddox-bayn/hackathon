package main

func BTreeLevelCount(root *TreeNode) int {
	// count to track data
	count := 0
	// base case to handle recursive traversal
	if root == nil {
		return count
	}
	
	// loop through the left up to parent
	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)
	count++
	heighestCount := leftHeight
	if rightHeight > heighestCount {
		heighestCount = rightHeight
	}
	return  heighestCount+1
}