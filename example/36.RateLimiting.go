package main

import (
	"fmt"
	"time"
)

// レート制限はリソース使用量を管理し、QoS を保つために重要だ。
// Go はゴルーチン、チャネル、tickers を使ってうまくレート制限をサポートする。
func main() {
	// まずはレートを制限する基本的なやり方を紹介する。
	// リクエストを受け取る量を制限したいとする。
	// このリクエストをチャネルに流し込む。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// チャネル limiter は200ミリ秒ごとに値を受信する。
	// このチャネルがレート制限のための制御役になる。
	limiter := time.Tick(200 + time.Millisecond)

	// limiter チャネルからの受信がブロックするのを利用して、200ミリ秒ごとにリクエストを受信する。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// リクエストの短期的なバーストを許容しながらも、
	// 長期的にはレート制限を守らせることもできる。
	// limiter にバッファを付ければいいのである。
	// チャネル burstyLimiter は3つまでのリクエストのバーストを許容する。
	burstyLimiter := make(chan time.Time, 3)

	// バーストを表す3つの要素をチャネルに送る。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 200ミリ秒ごとに、burstyLimiter に値を送信する。
	// ただし、要素数はチャネルの要素数は最大3である。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// ここで、5個リクエストが届いたとする。
	// そのうちはじめの3つは burstyLimiter のバースト耐性によって直ちに処理される。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	// プログラムを実行すると、はじめに送ったリクエストの一団は期待した通り200ミリ秒ごとに処理されたことがわかる。
	// その後に送られるリクエストのうち、はじめの3つはすぐに処理される。 これはレート制限がバーストを許容するからだ。
	// 最後に、残りの2つのリクエストが200ミリ秒ごとに処理される。
}
