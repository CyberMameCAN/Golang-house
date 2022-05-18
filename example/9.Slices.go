package main

import "fmt"

func main() {
	// スライスを作る時は make を使う
	s := make([]string, 3)	// 文字列が3つ入るスライスを作る（初期値はゼロ値）
	fmt.Println("emp", s)

	// 値の読み書き
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))	// スライスの長さ

	// 配列にあった基本的な操作に加えて、スライスにはより豊富な操作が可能である。
	// 例えば組み込みの append は一つかそれ以上の新たな値を含むスライスを返す。
	// ここで、新たな値を得るには、append の返り値を受け取る必要があることに注意する。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// スライスはコピー（copy）することもできる。
	// ここでは、s と同じ長さの空のスライス c を作り、s の内容を c にコピーしている。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// スライスを「スライス」する
	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]	// 5 未満
	fmt.Println("sl2", l)

	l = s[2:]	// 2 以上
	fmt.Println("sl3", l)

	// スライスの宣言と定義を、一行で行う
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// スライスを組み合わせて多次元のデータ構造を作れる。
	// 内側のスライスの長さは同じでなくてもよく、これは配列の場合とは異なる。
	twoD := make([][]int, 3)	// [[] [] []]
	fmt.Println(twoD)
	for i:=0; i<3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)	// [0] [1 2] [2 3 4]]
}