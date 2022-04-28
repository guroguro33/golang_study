package context

import (
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second) // 2秒待機
	fmt.Println("finish")
	ch <- "result"
}

func context() {
	ch := make(chan string)
	// 今回のテーマのcontextを作成
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch) // 作成したctxを渡す（特に使わないが）

	for {
		select { // for-selectでchが入ってくるまで待機している
		case <-ctx.Done(): // ctxが実行されたらこちらに入る
			fmt.Println(ctx.Err())
			return
		case <-ch:
			fmt.Println("success")
			return
		}
	}
}
