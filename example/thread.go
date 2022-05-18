/*
  第5章　並行プログラミング―ゴルーチンとチャネルを使いこなす
  https://gihyo.jp/dev/feature/01/go_4beginners/0005

  ゴルーチン：軽量な並行処理プログラミング
  チャネル：複数のゴルーチン間でデータのやり取りをする
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
func sampleSelectCase() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	for {
		// どれか１つcaseが実行されない限りは、select文はブロックする
		select {
		case c1 := <-ch1:
			// ch1からデータを読み出した時に実行
		case c2 := <-ch2:
			// ch2からデータを読み出した時に実行
		case ch2 <- "c":
			// ch2にデータを書き込んだ時に実行
		default:
			// caceが実行されなかった場合に実行
			// defaultの実行が終わるとselect文の処理が終わるため、select文がブロックされなくなる
		}
	}
}
*/

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string)

	for _, url := range urls {
		// wait.Add(1)

		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			// fmt.Println(url, res.Status)
			statusChan <- res.Status

			// wait.Done()
		}(url)
	}

	return statusChan
}

func main() {

	// wait := new(sync.WaitGroup)
	urls := []string{
		"https://keiba.to48.org/",
		"https://note.com/cybermamecan/",
		"https://egingerz.to48.org/",
	}

	statusChan := getStatus(urls)

	// main が終了しないように待ち合わせる
	// time.Sleep(time.Second)
	// wait.Wait()

	// ゴルーチンでstatusChanに値が書き込まれるまで、main()の中では値を読み出すことが出来ない。
	// ここではstatusChanの読み出しが３回完了するまで処理がブロックされる。
	// よってWait処理は必要ない。
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
