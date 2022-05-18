package main

import "fmt"

// デフォルトではチャネルにバッファは付いていない。
// すなわち、チャネルが送信（chan <-）を受け入れるのは、
// 受信（<- chan）の準備が既にできているときだけだ。
// バッファ付きチャネルは受信者がまだいなくても、一定の個数までなら送信を受け入れる。

func main() {
	// make を使って、2つまで値を溜められるバッファ付きチャネルを作る。
	messages := make(chan string, 2)

	// このチャネルにはバッファが付いているので、まだ受信者がいないにもかかわらずこれらの値を送れる。
	messages <- "buffered"
	messages <- "channel"

	// これらの値は後から読み出せる
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}