package main

import "fmt"

type person struct {
	name string
	age  int
}

// ローカル変数は関数のスコープを抜けても残るので、ローカル変数のポインタを返しても安全である。
func newPerson(name string) *person {
	// newPerson は名前受け取って新たな person を作る
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	// 構造体作成1
	fmt.Println(person{"Bob", 20})
	// 構造体を初期化するときに、フィールド名を指定してもよい
	fmt.Println(person{name: "Alice", age: 30})
	// 明記しないフィールドの値はゼロ値になる
	fmt.Println(person{name: "Fred"})
	// & を前に付けると構造体のポインタを作れる
	fmt.Println(&person{name: "Ann", age: 40})
	// コンストラクタ関数を使って構造体の作成をカプセル化するのは良くある書き方だ
	fmt.Println(newPerson("Jon"))

	// 構造体のフィールドにアクセスするにはドット演算子を使う
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// 構造体のポインタにもドット演算子を使える。
	// このとき、ポインタは自動的に参照を剥がされる。
	sp := &s
	fmt.Println(sp.age)
	// 構造体は可変である
	sp.age = 51
	fmt.Println(sp.age)
}