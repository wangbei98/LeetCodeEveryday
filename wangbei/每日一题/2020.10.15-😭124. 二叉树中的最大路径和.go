package main

import "math"

/*
方法1:
f(root) 计算从叶节点到root的最大路径
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxPathSum(root *TreeNode) int {
	maxPath := math.MinInt32
	// 计算从当前节点到叶子节点的最大路径
	var calMaxSingleSum func(root *TreeNode)int
	calMaxSingleSum = func(root *TreeNode)int{
		if root == nil{
			return 0
		}
		// 如果小于0，那这一支路线贡献为负，舍弃掉
		leftMaxSingleSum := max(calMaxSingleSum(root.Left),0)
		rightMaxSingleSum := max(calMaxSingleSum(root.Right),0)

		sum := leftMaxSingleSum + rightMaxSingleSum + root.Val
		maxPath = max(maxPath,sum)

		return root.Val + max(leftMaxSingleSum,rightMaxSingleSum)
	}
	_ = calMaxSingleSum(root)
	return maxPath
}

/*
方法二
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum2(root *TreeNode) int {
	var calPaths func(root *TreeNode)(int,int)
	calPaths = func(root *TreeNode)(int,int){
		if root == nil{
			return 0,math.MinInt32
		}
		// 计算以root为根的树的 最大单支路径和，最大路径和
		leftSingle,leftPath := calPaths(root.Left)
		rightSingle,rightPath := calPaths(root.Right)

		leftGain,rightGain := max(leftSingle,0),max(rightSingle,0)

		curSingle := root.Val + max(leftGain,rightGain)

		curPath := max(root.Val+leftGain+rightGain,max(leftPath,rightPath))
		return curSingle,curPath
	}
	_,res := calPaths(root)
	return res
}

func max(x,y int)int{
	if(x>y){
		return x
	}
	return y
}



func max(x,y int)int{
	if(x>y){
		return x
	}
	return y
}