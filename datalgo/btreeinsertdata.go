package main

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}
	if root == nil {
		return newNode
	}

	current := root
	var parent *TreeNode = nil
	for current != nil {
		if newNode.Data < current.Data {
			current.Left = newNode
		}
		if newNode.Data > current.Data {
			current.Left = newNode
		}
	}
	return root
}
