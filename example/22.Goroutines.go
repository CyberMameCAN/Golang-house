package main

import (
	"fmt"
	"time"
)

// ゴルーチンは軽量なスレッドである。

func f(from string) {
	for i:=0; i<3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 通常呼び出し
	f("direct")
	// go f(s) と書くと、この関数をゴルーチンの中で実行する。
	// こうすると、新たなゴルーチンが呼び出し側と平行に実行される
	go f("goroutine")
	// 無名関数のゴルーチンを起動することもできる
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 2つの関数呼び出しは、別々のゴルーチンにおいて非同期的に実行されている。
	// そのため処理はここまで抜けてくる。
	// ここで処理が終わるのを待つ（WaitGroup を使うとより確実だ）
	time.Sleep(time.Second)
	fmt.Println("done")

	// このプログラムを実行すると、ブロックする呼び出しの出力がまず表示され、
	// その後2つのゴルーチンの出力が入り混じって表示される。
	// こうなるのは Go のランタイムがゴルーチンを平行に実行しているためである。
}