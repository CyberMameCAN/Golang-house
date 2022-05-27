package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func RandomWait(ctx context.Context, i int) {
	// regionを始める
	defer trace.StartRegion(ctx, "randomWait").End()

	fmt.Printf("No.%d start\n", i+1)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("no.%d done\n", i+1)
}

func _main() {
	// タスクを定義
	ctx, task := trace.NewTask(context.Background(), "main")
	defer task.End()

	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			RandomWait(ctx, i)
		}(i)
	}
	wg.Wait()
}

func main() {
	// トレースを始める
	// 結果出力用のファイルもここで作成
	f, err := os.Create("tseq.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()

	_main()
}
