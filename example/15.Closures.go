package main

import "fmt"

// 無名関数
// 無名関数を使ってクロージャを作る

// intSeq は、intSeq の中で定義した無名関数を返す。
// 返される関数に変数 i を閉じ込めており、クロージャを作っている。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}