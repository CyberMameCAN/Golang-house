package main

import (
	"fmt"
	"context"
	"log"

	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"

	"google.golang.org/api/option"
	"google.golang.org/api/iterator"
	//"cloud.google.com/go/firestore"
	// "github.com/carlescere/scheduler"
	// "runtime"
)

// func createClient(ctx context.Context) *firestore.Client {
// 	//projectID := "serene-champion-305213"
// 	projectID := "/Users/toshi/Documents/Work/Golang/src/github.com/me/example/path/to/serviceAccount.json"

// 	client, err := firestore.NewClient(ctx, projectID)
// 	if err != nil {
// 		log.Fatal("Failed to create client: %v", err)
// 	}
// 	// defer client.Close()
// 	return client
// }

func main() {
	// scheduler.Every(5).Seconds().Run(printSuccess)
	// runtime.Goexit()
	// fmt.Println("Hello, world")

	// 各自のサーバで初期化する場合
	projPath := "/Users/toshi/Documents/Work/Golang/src/github.com/me/example/path/to/serviceAccount2.json"
	ctx := context.Background()
	//client := createClient(ctx)
	sa := option.WithCredentialsFile(projPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// 切断
	defer client.Close()	// defer: 関数の終了時に実行される処理を記述

	// 追加
	// _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
	// 	"first": "Ada",
	// 	"middle": "Mathison",
	// 	"last": "Lavelace",
	// 	"born": 1816,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed adding alovelace: %v", err)
	// }

	// データ読み取り
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

}

// func printSuccess() {
// 	fmt.Println("success!!\n")
// }