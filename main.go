package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// semaphoreのNewWeightedで同時に実行できる最大ゴルーチン数を設定する
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// TryAcquireでセマフォを1つ取得し、成功したらtrueを返す
	// 他のゴルーチンはセマフォを取得できず、falseが返り、待機ではなくreturnで終了してしまう
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("cloud not get lock")
		return
	}

	// // Acquireでセマフォを１つ取得する、今回は最大１のため、セマフォは0になり、他のゴルーチンは待機状態になる
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// セマフォをリリースして１に戻す。
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	// 空のcontextを生成
	ctx := context.TODO()
	// fmt.Println(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
