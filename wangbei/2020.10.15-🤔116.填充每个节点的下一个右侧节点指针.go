package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil{
		return root
	}
	queue := make([]*Node,0)
	queue = append(queue,root)
	for len(queue) > 0{
		length := len(queue)
		for i:=0;i<length-1;i++{
			queue[i].Next = queue[i+1]
		}
		queue[length-1].Next = nil
		for i:=0;i<length;i++{
			node := queue[0]
			queue = queue[1:]
			if node.Left!=nil{
				queue = append(queue,node.Left)
			}
			if node.Right != nil{
				queue = append(queue,node.Right)
			}
		}
	}
	return root
}