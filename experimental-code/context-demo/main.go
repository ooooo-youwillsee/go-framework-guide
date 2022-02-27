package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			fmt.Println("wait one second")
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("receive Done event: ", ctx.Err())
				return
			default:
				//fmt.Println("xxx")
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func main2() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	go func(ctx context.Context) {
		for {
			fmt.Println("wait one second")
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("receive Done event: ", ctx.Err())
				return
			default:
				//fmt.Println("xxx")
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(6*time.Second))

	go func(ctx context.Context) {
		for {
			fmt.Println("wait one second")
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("receive Done event: ", ctx.Err())
				return
			default:
				//fmt.Println("xxx")
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
