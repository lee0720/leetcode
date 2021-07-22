package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var drow = []int{-1, 1, 0, 0, -1, 1, -1, 1}
var dcol = []int{0, 0, -1, 1, -1, 1, 1, -1}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	ptr := 0
	for {
		ptr = nums[ptr]
		slow = nums[slow]
		if ptr == slow {
			break
		}
	}
	return ptr
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
