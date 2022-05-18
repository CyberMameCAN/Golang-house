package main

import (
	"fmt"
	"os"
)

// Defer を使って、ある関数を確実に後で呼ばせることができる。
// これは後片付けにしばしば使われる機能だ。
// 他の言語だと ensure や finally といったキーワードを使う場面で、Go では defer を使う。

// ファイルを作り、書き込み、最後にファイルを閉じたいとする。
// これを defer を使ってやってみよう。
func main() {
	// createFile でファイルオブジェクトを得た直後に、
	// ファイルを閉じるよう closeFile を defer しておく。
	// 直近の関数（main）の最後、つまり writeFile が終わるときに、defer した関数が実行される。
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// defer を使ったときにも、ファイルを閉じた際のエラーチェックを忘れないように注意しよう。
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
