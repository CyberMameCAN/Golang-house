# Golang 覚え書き
go1.17以降、**go get**でのパッケージインストールが非推奨、代わりに**go install**を使うように。

go get -u オプションの**-u**は新しいリリースまたはパッチリリースが利用可能な場合、パッケージとその依存パッケージをネットワークから更新する。使用しているパッケージを更新する際も使うことができる。Updateのuか。

**importのアンダースコア** インポートするパッケージの関数は直接は使わないけど、依存関係によってインポートの必要があり、init()だけを実行する？

$ go mod tidy  使われていない依存モジュールを削除する


## Semaphore(セマフォ)
同時実行を制御するための仕組み・手法

## ガーベイジコレクション

コンピュータプログラムが動的に確保したメモリ領域のうち、不要になった領域を自動的に解放する機能。

## JWT (JSON Web Token)

トークン内に情報を保持し、それを署名、暗号化することができる技術  
つまり、署名を検証する事でユーザの認証が可能であり、内容に認可の情報を含める事で認証から認可まで行うことが可能になる。

### JWK (JSON Web Key)

暗号鍵をJSONを使用して表現するための方法


## ティッカー

Linuxでいうcronのような機能

## 慣用

### アセンブリ

$ otool -t -v -V mybinary > result.txt  <- MacOS
$ readelf <- Linux
$ go tool objdump main > objdump.txt

### 型の調べ方
	reflect.TypeOf(xxx)
	
### メソッド、構造体
- 名前の先頭が大文字 -> 公開メソッド
- 　　　　　　小文字 -> 非公開メソッド(privateになる)
- メソッドチェーンはあまり使わない。なぜならerrorを返すのでif文で処理するため

### パッケージ管理ツール

- dep
- vgo

とりあえずdepを使ってみた。(vgoは公式ツール)

#### dep

- dev version
- dev init

### Windows用にコンパイル（クロスコンパイル）

`# GOOS=windows GOARCH=amd64 go build main.go`

Linux, MacOSX に**file**というコマンドがあるのを知った。
ファイルのメタ情報を表示できる。

クロスコンパイル時に指定する環境変数(GOOS, GOARCH)へ指定できる値の確認コマンド
`# go tool dist list`

## Webフレームワーク
- gin
- echo
- gorilla/mux
- Beego (これを使う)
- FastAPI (これも使いたい)
- [Fiber](https://docs.gofiber.io/extra/benchmarks)

## 変数定義
### 普通の変数
1. var <変数名> <型>
2. <変数名> := <初期値>
3. var <変数名> <型> = <初期値>

### 配列宣言
	var a [10]int
	s := []int{7, 2, 8, -9, 4, 0}
	
配列の長さを変えることはできない。

### スライス
Numpy(Python)のスライスと使い方は同じ？ではなく、大きさを変えられる配列

#### 定義例
	var (
	    NotPV []string = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
	)


### 定数
	const Pi = 3.14
	const big = 0xFFFFFF
	
### 型を調べる
	reflect.TypeOf(XXX)


## 関数定義
### 2つの引数、2つの戻り値の場合
	func myFunc(name string, age uint) (bool, error) {
	    var find bool
	    var result error
	    return find, error
	}
### 値だけど内部はポインタっぽいのはそのまま渡す
	slice, map, chan, func,等
	
## メソッド
**レシーバ**と呼称。関数の前にレシーバを記述する形でメソッドを宣言できる。

	func (p Calc) Add() {
	    return p.atai1 + p.atai2
	}
	
	func (レシーバ 型) 関数名(引数){
	  // 処理
	}
	
### ポインタレシーバ
Calc型構造体をレシーバpとして受け取る

	func (p *Calc) Add() {
	    return p.atai1 + p.atai2
	}


## Golang特有の if の書き方
### if <変数> := <処理>; <変数がtrue or falseにより処理>
	func myTest() bool {
		...
	}
	
	func main() {
		if res := muTest(); !res {
			// この場合はFalse
		}
	}

## 構造体
	type 名称 型 {
		名前 型
		名前 型
		名前 型
	}

### 要素にアクセス
ドットでアクセスする。

	var 変数 名称
	名称.名前

### ※足し算、掛け算などに注意あり

## 関数その他
- import( os, fmt), fmt.Fprintf(os.Stterr, "err '%s'", err)

## 今後は
- interface
- channel
- context

などを勉強すると良い。
goroutineは並列処理を実装

ポインタ型はあるがポインタ演算は無い。

[A Tour of Go](https://go-tour-jp.appspot.com/list)をまずは１周。

## interface
interface型は、メソッドのシグニチャ(用法？)の集まり

	type Abser interface {
		Abs() float64
	}

## Goroutines

	go f(x, y, z)
	
と書けば新しい**goroutine (並列処理)**が実行される

## Emacs
### No Windowモードで起動

`$ emacs -nW`

### コマンド実行
`M-x`

## REST

分散システムにおける複数のソフトウェアを連携させるのに適した設計原則の集合、考え方のこと

* URIではリソースの指定を行う
* HTTPメソッド(POST(追加), PUT(更新), GET(取得), DELETE)で設計する。
* ステートレスであること（セッションといった状態を管理しない）
* HTML, XML, JSONなどのハイパーメディアの使用(やり取りするデータのこと)

### gRPC リモートプロキシコール

Golang界隈では有名

Protcol Buffer
  gRPCとを実現するためのシリアライズフォーマット。.protoファイルに記述する。
  
## github開発のやりかた

1. [github.com] 新しくリポジトリ作成(Add .gitignore Goを選ぶ、ライセンスなど追加)
2. [Mac, PC] git clone [CodeのSSHをコピー ex) git@github.com:CyberMameCAN/リポジトリ名.git]
3. cd リポジトリ名
4. .gitignore編集（vender/ を有効に、app/アプリ名を追加）
5. Push all the code with git: (最後のorigin master は mainに変更になってる)

	$ git add Dockerfile*
	$ git add src
	$ git add .gitignore
	$ git commit -m "initial commit"
	$ git push origin main

6. $ git status 次に何するか確認

### gitのコマンド
#### 最初だけ
1. git init
2. git remote add origin <URL>

#### 基本的開発の流れ
3. git swich -c <ブランチ名>
4. git add .  変更内容をステージングに追加
5. git commit -m "メッセージ"
6. git push origin <ブランチ名>  <- main 


7. git swich <ブランチ名>  ブランチを切り替える
8. git pull origin <ブランチ名> GitHubの変更内容を取り込む

#### よく使うコマンド
9. git status 変更したファイルを確認
10. git diff 変更したファイル内容を確認
11. git branch ブランチの一覧表示

## Go routine

メインゴールーチンが終わったら、他スレッド？のゴールーチンの終了を待たずにプログラム全体が終わる。

### チャネル(Channel)
特定の型の値を送信・受信する事で（異なるゴールーチンで）並行に実行している関数がやり取りする機構を手起居している。
チャネル演算子 **<-**

	ch <- v // vをチャネルchへ送信する
	v := <-ch // chから受信した変数をvへ割り当てる
	
#### チャネルの生成

	ch := make(chan int)

通常、片方が準備できるまで送受信はブロックされる。これにより、明確なロックや条件変数がなくても、goroutineの同期を可能にする。

	package main
	
	import "fmt"
	
	func sum(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum // send sum to c
	}
	
	func main() {
		s := []int{7, 2, 8, -9, 4, 0}
		
		c := make(chan int)
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:, c)
		x, y := <-c, <- c // receive from c
		
		fmt.Println(x, y, x+y)
	}

## Redis

Key-Valueで値が保存されるNoSQLデータベース  
常にメモリにデータが乗るため、メモリ容量しか扱えない、ゆえにメモリの消費が激しい。


## プログラミング基本構文

### 変数と定数
### データ型
### 関数
### 制御構文

## Heroku + Docker

	$ mkdir projectname
	$ cd projectname
	$ git init
	$ git add .
	$ git commit -m "My first commit"
	
	$ heroku create # プロジェクトができる
	$ git remote -v  # できているか確認
	$ heroku git:remote -a [thawing-peak-40393みたいなの]
		$ git remote rename heroku [heroku-stagingとか] # リモート名heroku の変更
		
	$ touch heroku.yml
	$ git add heroku.yml
	$ git commit -m "Add heroku.yml"
	$ heroku stack:set container
	