package main

import (
	"fmt"
)

func main() {
	// generator
	generator := func(done <-chan interface{}, integers ...int) <-chan int { // ...int 複数のintを渡せる
		intStream := make(chan int, len(integers)) // チャンネルを作るときはmakeを使う, len(integers)はバッファ
		go func() {                                // 無名関数の定義
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	// multiply
	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int) // バッファーを用意していない（相手が受け取ってくれるまでブロック）
		go func() {                        // goルーチンの生成
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:

				}
			}
		}()
		return multipliedStream
	}

	// add
	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int) // バッファーを用意していない（相手が受け取ってくれるまでブロック）
		go func() {                   // goルーチンの生成
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:

				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{}) // interface{} 全て型
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
