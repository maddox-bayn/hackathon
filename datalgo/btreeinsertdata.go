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
	// handle cases where struct exist but has no  value
	if root.Data == "" && root.Left == nil && root.Right == nil {
		root.Data = data
		return root
	}
	current := root
	var lastValideNode *TreeNode = nil
	
	for current != nil {
		lastValideNode = current
		if data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	// make connection and between parent and child node
	newNode.Parent = lastValideNode
	if data < lastValideNode.Data {
		 lastValideNode.Left = newNode
	} else {
		lastValideNode.Right = newNode 
	}
	return root
}
