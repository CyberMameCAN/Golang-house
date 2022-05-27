package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// クライアント生成
	client := &http.Client{Timeout: time.Duration(30) * time.Second}
	// リクエスト送信
	req, err := http.NewRequest("GET", "http://localhost:9880/basic_auth", nil)
	if err != nil {
		log.Fatal(err)
	}
	// 認証情報をセット
	req.SetBasicAuth("client_id", "client_secret")
	// リクエスト実行
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
