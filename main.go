package main

import (
	"fmt"
	"sort"
)

var drow = []int{-1, 1, 0, 0, -1, 1, -1, 1}
var dcol = []int{0, 0, -1, 1, -1, 1, 1, -1}

func main() {
	merge([]int{1, 3, 4, 0, 0, 0}, 3, []int{2, 6, 7}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	cur := m + n - 1
	tail := 0
	for p1 >= 0 || p2 >= 0 {
		if p1 == -1 {
			tail = nums2[p2]
			p2--
		} else if p2 == -1 {
			tail = nums1[p1]
			p1--
		} else if nums1[p1] >= nums2[p2] {
			tail = nums1[p1]
			p1--
		} else if nums1[p1] < nums2[p2] {
			tail = nums2[p2]
			p2--
		}

		nums1[cur] = tail
		cur --
	}
	fmt.Println(nums1)
}

func inorderSuccessor(root *TreeNode, p *TreeNode) (t *TreeNode) {
	var (
		flag bool
		ans  *TreeNode
		f    func(*TreeNode)
	)
	f = func(r *TreeNode) {
		if r != nil && ans == nil {
			f(r.Left)
			if r == p {
				flag = true
			} else if flag {
				ans, flag = r, false
			}
			f(r.Right)
		}
	}
	f(root)
	return ans
}

func inorderSuccessorV1(root *TreeNode, p *TreeNode) (t *TreeNode) {
	var res *[]int
	var inorder func(root *TreeNode, r *[]int)
	var inorder1 func(root *TreeNode, target int, t *TreeNode)
	inorder = func(root *TreeNode, r *[]int) {
		if root == nil {
			return
		}
		inorder(root.Left, r)
		*r = append(*r, root.Val)
		inorder(root.Right, r)
	}

	inorder(root, res)
	var val int
	for i := range *res {
		if (*res)[i] == p.Val {
			if i+1 <= len(*res) {
				val = (*res)[i+1]
			} else {
				return nil
			}
		}
	}

	inorder1 = func(root *TreeNode, target int, t *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == target {
			t = root
		}
		inorder1(root.Left, target, t)
		inorder1(root.Right, target, t)

	}
	inorder1(root, val, t)
	return t

}

func pondSizes(land [][]int) (res []int) {

	maps := make([][]bool, len(land))
	for i := range land {
		maps[i] = make([]bool, len(land[i]))
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if land[i][j] == 0 && maps[i][j] == false {
				count := 1
				maps[i][j] = true
				dfs(i, j, &maps, land, &count)
				res = append(res, count)

			}

		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func dfs(row, col int, maps *[][]bool, land [][]int, count *int) {
	n := len(land)
	m := len(land[0])
	for i := range dcol {
		cCol := col + dcol[i]
		cRow := row + drow[i]
		if cRow < 0 || cRow >= n {
			continue
		}

		if cCol < 0 || cCol >= m {
			continue
		}

		if (*maps)[cRow][cCol] == true {
			continue
		}

		if land[cRow][cCol] != 0 {
			continue
		}
		(*maps)[cRow][cCol] = true
		(*count) += 1
		dfs(cRow, cCol, maps, land, count)

	}

}

// 给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。

// 示例:
// 给定如下二叉树，以及目标和 sum = 22，

//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	left, right := 0, 0
	count := 0
	preOrder(root, 0, sum, &count)
	if root.Left != nil {
		left = pathSum(root.Left, sum)
	}
	if root.Right != nil {
		right = pathSum(root.Right, sum)
	}

	return count + left + right
}

func preOrder(root *TreeNode, cur, sum int, count *int) {
	if root == nil {
		return
	}
	cur = cur + root.Val
	if cur == sum {
		*count = *count + 1

	}
	preOrder(root.Left, cur, sum, count)
	preOrder(root.Right, cur, sum, count)
}
