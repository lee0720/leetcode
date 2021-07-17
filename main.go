package main

import (
	"fmt"
	"sort"
)

var drow = []int{-1, 1, 0, 0, -1, 1, -1, 1}
var dcol = []int{0, 0, -1, 1, -1, 1, 1, -1}

func main() {
	arr := [][]int{
		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}

	fmt.Println(pondSizes(arr))
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
