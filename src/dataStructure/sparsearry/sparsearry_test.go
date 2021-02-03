package  sparsearry

import (
	"fmt"
	"testing"
)

type ValNode struct {
	row int
	col int
	val int
}

// 稀疏数组
func TestSparseArray(t *testing.T) {

	// 原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 蓝子

	// 输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 转成稀疏数组
	// 遍历，如果发现有一个数组值不为0，创建一个node结构体，放入结构体
	// 标准稀疏数组有记录二维数组的规模（行、列、值）
	sparseArr := []ValNode{}
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			// 创建一个valnode值节点
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	// 恢复稀疏数组
	chessMap2 := [11][11]int{}

	// 遍历稀疏数组
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	// 恢复后数据
	fmt.Println("恢复后的原始数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
