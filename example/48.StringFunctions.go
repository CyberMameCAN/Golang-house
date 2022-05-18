package main

import (
	"fmt"
	s "strings"
)

// 標準ライブラリの strings パッケージには文字列を扱うための便利な関数が含まれている。
// ここではこのパッケージをどういう感じで使えるのかがわかる例を紹介する。

// fmt.Println を繰り返し使うので、短い別名を与える。
var p = fmt.Println

func main() {
	// strings の関数のサンプルを出そう。
	// これらの関数はパッケージの関数であって、メソッドではないので、
	// 第一引数として対象となる文字列を渡す。
	// strings のドキュメントを見れば他にも色々な関数があることがわかる。
	p("Contains: ", s.Contains("test", "es"))
	p("Count:    ", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:    ", s.Index("test", "e"))
	p("Join:     ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:   ", s.Repeat("a", 5))
	p("Replace:  ", s.Replace("foo", "o", "0", -1))
	p("Replace:  ", s.Replace("foo", "o", "0", 1))
	p("Split:    ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:  ", s.ToLower("TEST"))
	p("ToUpper:  ", s.ToUpper("test"))
	p()

	// strings パッケージの一部ではないが、文字列の長さを求める機能、
	// インデックスを指定してバイトを読み出す機能もここで紹介しておくべきだろう。
	p("Len: ", len("hello"))
	p("Char: ", "hello"[1])
}
