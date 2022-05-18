// p は読み込んだ内容を一時的に入れておくバッファ
// func Read(p []byte) (n int, err error)

// Golang でメモリを確保するには make を使うと良い
// buffer := make([]byte 1024)
// size は実際に読み込んだバイト数
// size, err := r.Read(buffer)

// io.Reader の補助関数

// 全て読み込む
// buffer, err := ioutil.ReadAll(reader)

// 4バイト読み込めないとエラー
// buffer := make([]byte, 4)
// size, err := io.ReadFull(reader, buffer)

// コピーの補助関数

// io.Reader から io.Writer にそのままデータを渡したい場合に使うのがコピー系の補助関数

// writeSize, err := io.Copy(writer, reader)	// 全てコピー
// writeSize, err := io.CopyN(writer, reader, size)	// 指定したサイズだけコピー

// デフォルトではio.COpyは32KBのバッファを内部に持つ
// buffer := make([]byte, 8 * 1024)	// 8KBのバッファを使う
// io.CopyBuffer(writer, reader, buffer)

// io.Reader を満たす構造体で、よく使うもの

package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func readConsole() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}

func fileCopy() {
	file, err := os.Open("http.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

func internetConnect() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

func internetResponsPers() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダを表示
	fmt.Println(res.Header)
	// ボディを表示(最後はClose処理をすること)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func makeFileCopy() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, oldFile)
}

func makeGoodyFile() {
	fp, err := os.Create("goodfile.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// io.Copy(fp, strings.NewReader("私の名前は竹下敏也です。"))
	io.CopyN(fp, rand.Reader, 1024)
}

func main() {
	// readConsole()
	// fileCopy()
	// internetConnect()
	// internetResponsPers()
	// makeFileCopy()
	makeGoodyFile()
}
