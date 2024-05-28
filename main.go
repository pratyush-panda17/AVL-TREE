package main

import "fmt"

type Tnode struct { //Creating recursive structure for a Tree node
	val    int
	left   *Tnode
	right  *Tnode
	height int
}

type Tree struct { // An AVL tree structure that always points to the root node
	root *Tnode
}

func createNode(v int) *Tnode { //Initializing a node
	return &Tnode{val: v, left: nil, right: nil, height: 1}

}

func (bst *Tree) GetHeight() int { //Function to get height of bst tree
	return getHeight(bst.root)
}

func getHeight(root *Tnode) int {
	if root == nil {
		return 0
	}
	return root.height
}

func setHeight(root *Tnode) int { // Updates height of a tree node after insertion or rotation
	a := getHeight(root.left)
	b := getHeight(root.right)
	return max(a, b) + 1
}

func getBf(root *Tnode) int { // Gets balance factor of node to determine imbalances in right or left subtree
	if root == nil {
		return 0
	}
	return getHeight(root.left) - getHeight(root.right)
}

func leftRotate(root *Tnode) *Tnode { // Performs left rotation on given node
	right_child := root.right
	left_child := right_child.left

	right_child.left = root
	root.right = left_child

	root.height = setHeight(root)
	right_child.height = setHeight(right_child)

	return right_child
}

func rightRotate(root *Tnode) *Tnode { // Performs right rotation on given node
	left_child := root.left
	right_child := left_child.right

	left_child.right = root
	root.left = right_child

	root.height = setHeight(root)
	left_child.height = setHeight(left_child)

	return left_child
}

func (bst *Tree) InsertNode(val int) { // Inserting a value into the tree
	bst.root = insertNode(bst.root, val)
}

func insertNode(root *Tnode, val int) *Tnode {
	if root == nil {
		return createNode(val)
	} else if val < root.val {
		root.left = insertNode(root.left, val)
		if getBf(root) > 1 { // Tree may have a left imbalance after inserting value into left subtree
			if val < root.left.val { // Determines if the inserted node caused a LL imbalance or an LR imbalance
				root = rightRotate(root)
			} else {
				root.left = leftRotate(root.left)
				root = rightRotate(root)
			}
		}
	} else {
		root.right = insertNode(root.right, val)
		if getBf(root) < -1 { // Tree may have a right imbalance after inserting value into right subtree
			if val > root.right.val { // Determines if the inserted node caused a RR imbalance or an RL imbalance
				root = leftRotate(root)
			} else {
				root.right = rightRotate(root.right)
				root = leftRotate(root)
			}
		}
	}
	root.height = setHeight(root) // Recursively updates height of every node after insertion
	return root
}

func (bst *Tree) DeleteNode(val int) { // Function for deleting a node in a bst
	bst.root = deleteNode(bst.root, val)
}

func deleteNode(root *Tnode, val int) *Tnode {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = deleteNode(root.left, val)
		if getBf(root) < -1 { // Tree may have a right imbalance after deleting value from left subtree
			if val > root.right.val { // Determines if the deleted node caused a RR imbalance or an RL imbalance
				root = leftRotate(root)
			} else {
				root.right = rightRotate(root.right)
				root = leftRotate(root)
			}
		}
	}
	if val > root.val {
		root.right = deleteNode(root.right, val)
		if getBf(root) > 1 { // Tree may have a left imbalance after deleting value from right subtree
			if val < root.left.val { // Determines if the deleted node caused a LL imbalance or an LR imbalance
				root = rightRotate(root)
			} else {
				root.left = leftRotate(root.left)
				root = rightRotate(root)
			}
		}
	} else { // If left or right node is nil it implies tree has only one or two values
		if root.left == nil {
			temp := root.right
			root = nil
			return temp
		} else if root.right == nil {
			temp := root.left
			root = nil
			return temp
		}
		temp := findSmallest(root.right)              // Updating value of root node to the smallest value from right sub tree
		root.val = temp.val                           // so that bst tree property is maintained.
		root.right = deleteNode(root.right, temp.val) // Deleting smallest node from right-sub tree
		if getBf(root) > 1 {                          // Tree may have a left imbalance after deleting value from right subtree
			if val < root.left.val {
				root = rightRotate(root)
			} else {
				root.left = leftRotate(root.left)
				root = rightRotate(root)
			}
		}
	}
	return root
}

func findSmallest(root *Tnode) *Tnode { // Function for finding the smallest node in a particular sub-tree
	for root.left != nil {
		root = root.left
	}
	return root
}

func (bst *Tree) Search(val int) bool { //Function that confirms whether a value exists in the tree
	return search(bst.root, val)
}

func search(root *Tnode, val int) bool {
	if root == nil {
		return false
	} else if root.val == val {
		return true
	} else if root.val < val {
		return search(root.right, val)
	}
	return search(root.left, val)
}

func (bst *Tree) InOrder() { // Prints in-order traversal of tree in a single line
	inOrder(bst.root)
}

func inOrder(root *Tnode) {
	if root != nil {
		inOrder(root.left)
		fmt.Printf("%d ", root.val)
		inOrder(root.right)
	}
}

func (bst *Tree) PreOrder() { // Prints pre-order traversal of tree in a single line
	preOrder(bst.root)
}

func preOrder(root *Tnode) {
	if root != nil {
		fmt.Printf("%d ", root.val)
		preOrder(root.left)
		preOrder(root.right)
	}
}

func (bst *Tree) PostOrder() { // Prints post-order traversal of tree in a single line
	postOrder(bst.root)
}

func postOrder(root *Tnode) {
	if root != nil {
		postOrder(root.left)
		postOrder(root.right)
		fmt.Printf("%d ", root.val)
	}
}

func main() {

}
