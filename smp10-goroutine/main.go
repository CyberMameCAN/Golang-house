package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	// fmt.Printf("Today's your lucky number is %d!\n", num)
	c <- num // チャネルに送信
}

func main() {
	fmt.Println("what is today's lucky number?")

	c := make(chan int) // チャネルを作成
	go getLuckyNum(c)

	num := <-c // チャネルから返り値を受信

	fmt.Printf("Today's your lucky number is %d!\n", num)
	close(c)

	// time.Sleep(time.Second * 5) // これをコメントアウトしたらgo routineの処理が終わる前に終了してしまう。

	// 別のゴールーチンが終わるのを待って終了したい場合
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	getLuckyNum()
	// }()

	// wg.Wait()
	// チャネルを使うと、メインゴールーチンはcを受信するまでブロックされるので、別のゴールーチンが終わることはない。
}
