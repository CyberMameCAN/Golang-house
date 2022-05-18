package main

import "os"

// panic は予期せぬ問題が起きたことを示すために使う。
// これは普通は起きるはずがない問題が起きたり、
// うまく対処できない問題が起きた場合に、
// なるべく早くプログラムを停止させるために使う。

func main() {
	// このウェブサイトでは、予期せぬエラーが発生した場合に panic を使う。
	// この行だけが、意図的に panic を呼んでいる箇所である。
	panic("a problem")

	// 関数がエラーを返し、そのエラーにどう対処すべきがわからない（または対処したくない）場合に、
	// panic を呼んでプログラムを中断するのは、panic の典型的な使い方である。
	// この例では、ファイルを作るときにエラーが発生したら、panic を呼んでいる。
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	// 例外を使って多くのエラーを処理する言語とは違って、
	// Go ではエラーの有無を返り値として返すのが普通であることに注意する。

	// err := recover()
	// deferの中にrecover()を書くことでパニックで発生したエラーの処理を実施してから関数を抜けることができる。
}
