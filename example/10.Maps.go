package main

import "fmt"

func main() {
	// マップを作るには make を使う
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	// delete はマップからキーと値の組を削除する
	delete(m, "k2")
	fmt.Println("map:", m)

	// オプションの2つ目の返り値は、マップにキーが含まれるかどうかを表す真偽値である。
	// これはキーが入っていない場合と、そのキーの値としてゼロ値（0 や "" など）が入っている場合とを
	// 区別するために使える。 ここでは、値はいらないので、ブランク識別子 _ を使って、無視している。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// マップの宣言と初期化を、１行で行う
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map:", n) // map[bar:2 foo:1]
}
