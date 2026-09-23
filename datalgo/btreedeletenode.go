package main

// BTreeDeleteNode deletes a specific node while preserving BST properties
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	// BASE CASE: Node not found or tree is empty
	if root == nil {
		return nil
	}

	// 1. Search phase: Drill down the tree to locate the target node
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
		return root
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
		return root
	}

	// 2. Deletion phase: 'root' matches 'node' here!
	
	// Case 1 & 2: No left child (covers leaf node and right-only child cases)
	if root.Left == nil {
		return root.Right // Replaces root with its right child/nil
	}
	// Case 2: No right child (covers left-only child case)
	if root.Right == nil {
		return root.Left // Replaces root with its left child
	}

	// Case 3: Node has two children
	// Find the smallest node in the right subtree (In-order Successor)
	successor := BTreeMin(root.Right)
	
	// Swap the data values
	root.Data = successor.Data
	
	// Delete the old successor node from the right subtree
	root.Right = BTreeDeleteNode(root.Right, successor)

	return root
}
