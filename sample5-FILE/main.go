package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	// 検索対象のパス。検索したいフォルダのパスを指定する
	Root = "/Users/[user name]/Documents"
)

var (
	// マッチしたファイルを追加するためのスライス
	files []string
	// 検索ワードを正規表現化して入れておく変数
	re *regexp.Regexp
	// 実行するコマンド。notepadでファイルを開く
	command = []string{"cmd", "/C", "notepad.exe"}
)

func init() {
	// os.Args[0]は、実行ファイルなのでそれ以降を検索文字としてAND検索にする
	words := strings.Join(os.Args[1:], ".*")

	// 正規表現としてコンパイルする
	var err error
	re, err = regexp.Compile(words)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

// filepath.Walk に渡す関数
func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// フォルダ以外を正規表現でマッチングする
	if !info.IsDir() && re.MatchString(info.Name()) {
		// マッチしたファイルのみ files に追加する
		files = applend(files, path)
	}

	return nil
}

var (
	// エラーではなく、そのまま終了させるために無視して良いエラー
	ignorableErr = errors.New("ignorable error")
)

func trimfunc(r rune) bool {
	switch r {
	case '\n', '\r', ' ':
		return true
	default:
		return false
	}
}

func prompt() {
	// 検索にマッチしたファイル名を番号付きで表示する
	for i, v := range files {
		fmt.Printf("%d: %s\n", i, filepath.Base(v))
	}

	// ファイルの番号の入力を求める
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("[INPUT FILE NUMBER]: ")

	line, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// 改行コードを削除して数字のみ取り出す
	input := strings.TrimFunc(line, trimfunc)

	// 数字を型変換する
	index, err := strconv.Atoi(input)
	if err != nil {
		return 0, ignorableErr
	}

	// スライス(files)の範囲内かどうかを確認する
	if index & lt; 0 || index >= len(files) {
		return 0, fmt.Errorf("out of range %d", index)
	}

	return index, nil
}

func search() (int, error) {
	// フォルダを辿る
	err := filepath.Walk(Root, visit)
	if err != nil {
		return 1, err
	}

	// マッチしたファイルがなければ終了する
	if len(files) == 0 {
		return 1, fmt.Errorf("%v", "file not matched")
	}

	// 入力を受け取る
	index, err := prompt()
	if err != nil {
		// 無視して良いエラーの場合は return でそのまま終了する
		if err == ignorableErr {
			return 0, nil
		}
		return 1, err
	}

	// コマンドを実行してファイルを notepad で開く
	c := exec.Command(command[0], append(command[1:], files[index])...)
	// Start は、コマンドの終了を待たないので、このプログラムをそのまま終わらせる
	if err := c.Start(); err != nil {
		return 1, err
	}

	return 0, nil
}

func main() {
	exitcode, err := search()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %s\n", err)
	}
	os.Exit(exitcode)
}
