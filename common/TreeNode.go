package common

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeFromSlice(data []interface{}) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	root := &TreeNode{Val: data[0].(int)}
	queue := list.New()
	queue.PushBack(root)

	for i := 1; i < len(data); i += 2 {
		currentNode := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())

		leftVal := data[i]
		rightVal := data[i+1]

		if leftVal != nil {
			currentNode.Left = &TreeNode{Val: leftVal.(int)}
			queue.PushBack(currentNode.Left)
		}

		if rightVal != nil {
			currentNode.Right = &TreeNode{Val: rightVal.(int)}
			queue.PushBack(currentNode.Right)
		}

	}

	return root
}

func (t *TreeNode) PreOrderRecur() {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	t.Left.PreOrderRecur()
	t.Right.PreOrderRecur()
}

func (t *TreeNode) InOrderRecur() {
	if t == nil {
		return
	}
	t.Left.InOrderRecur()
	fmt.Println(t.Val)
	t.Right.InOrderRecur()
}

func (t *TreeNode) PostOrderRecur() {
	if t == nil {
		return
	}
	t.Left.PostOrderRecur()
	t.Right.PostOrderRecur()
	fmt.Println(t.Val)
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		currentNode := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())

		fmt.Print(currentNode.Val, " ")

		if currentNode.Left != nil {
			queue.PushBack(currentNode.Left)
		}

		if currentNode.Right != nil {
			queue.PushBack(currentNode.Right)
		}
	}
}
