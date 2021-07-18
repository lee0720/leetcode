package main

import (
	"fmt"
	"sort"
	"strconv"
)

var drow = []int{-1, 1, 0, 0, -1, 1, -1, 1}
var dcol = []int{0, 0, -1, 1, -1, 1, 1, -1}

func main() {

	fmt.Println(trulyMostPopular([]string{"a(10)", "c(13)"}, []string{"(a,b)", "(c,d)", "(b,c)"}))
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

func trulyMostPopular(names []string, synonyms []string) []string {
	freq := map[string]int{}
	uf_set := map[string]string{}
	find := func(name string) string {
		if _, exits := uf_set[name]; !exits {
			return ""
		}
		for uf_set[name] != name {
			name = uf_set[name]
		}
		return name
	}

	union := func(name1, name2 string) {
		set1, set2 := find(name1), find(name2)
		if set1 != "" && set2 != "" && set1 != set2 {
			if set1 < set2 {
				uf_set[set2] = set1
				freq[set1] += freq[set2]
				delete(freq, set2)
			} else {
				uf_set[set1] = set2
				freq[set2] += freq[set1]
				delete(freq, set1)
			}
		}
	}

	for _, name_freq := range names {
		end := 0
		for name_freq[end] != '(' {
			end++
		}
		name := name_freq[:end]
		uf_set[name] = name
		freq[name], _ = strconv.Atoi(name_freq[end+1 : len(name_freq)-1])
	}

	for _, syn := range synonyms {
		end := 0
		for syn[end] != ',' {
			end++
		}
		name1 := syn[1:end]
		name2 := syn[end+1 : len(syn)-1]
		if _, exits := uf_set[name1]; !exits {
			uf_set[name1] = name1
		}
		if _, exits := uf_set[name2]; !exits {
			uf_set[name2] = name2
		}
		union(name1, name2)

	}

	res := []string{}
	for name := range freq {
		if freq[name] != 0 {
			res = append(res, name+"("+strconv.Itoa(freq[name])+")")
		}
	}
	return res
}

func findCircleNum(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(x, y int)
	union = func(x, y int) {
		parent[find(x)] = find(y)
	}

	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	for i := range parent {
		if i == parent[i] {
			ans++
		}
	}
	return ans

}
