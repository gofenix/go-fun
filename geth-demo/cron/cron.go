package main

import (
	"time"
	"log"
)

func main() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			time := <-ticker.C
			log.Println("定时器====>", time.String())
		}
	}()

	// 等待1分钟，观察输出结果
	time.Sleep(1*time.Minute)
}
