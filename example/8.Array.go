package main

import "fmt"

// Go ではここの配列より、スライスの方が使われる

func main() {
	// a は int 型の値5個からなる配列
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// このように書けば、配列の宣言と初期化を一行で済ませられる
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 多次元配列
	var twoD [2][3]int
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}