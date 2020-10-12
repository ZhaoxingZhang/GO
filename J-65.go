package main

 type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	 }

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B==nil||A==nil {
		return false
	}
	var result bool
	if A.Val==B.Val {
		result = doesT1HasT2(A,B)
	}
	if result==false && A.Left!=nil {
		result = isSubStructure(A.Left,B)
	}
	if result==false && A.Right!=nil {
		result = isSubStructure(A.Right,B)
	}
	return result
}
func doesT1HasT2(A *TreeNode, B *TreeNode) bool {
	if B==nil {
		return true
	}
	if A==nil {
		return false
	}
	if A.Val!=B.Val {
		return false
	}
	return doesT1HasT2(A.Left,B.Left) && doesT1HasT2(A.Right,B.Right)
}

