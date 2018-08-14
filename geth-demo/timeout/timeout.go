package main

import (
	"time"
	"log"
)

func main() {
	Test1()
	//Test2()
	//Test3()
	//Test4()
}

// 不能正常运行
func Test1() {
	isTimeout := false
	data := make(chan string)

	go func() {
		for {
			if isTimeout {
				log.Println("收到超时请求，关闭生产者协程")
				return
			} else {
				time.Sleep(1 * time.Second)
				data <- "生产数据"
			}

		}
	}()

	// 设置消费者超时时间
	after := time.After(5 * time.Second)
	go func() {

		for {
			select {
			case x := <-data:
				log.Println("从data通道拿数据：", x)

			case <-after:
				log.Println("消费者超时，不再消费")
				isTimeout = true
				return
			}
		}
	}()

	defer close(data)

	time.Sleep(60 * time.Second)
}

// isTimeout不能用变量的形式，因为此时是从data里面拿数据，如果生产者不生产的话，data通道没有数据，消费者就会阻塞。
func Test2() {
	isTimeout := make(chan bool)
	data := make(chan string) // 这里

	after := time.After(5 * time.Second)
	go func() {
		for {
			select {
			case <-after:
				log.Println("生产者超时，不再生产")
				isTimeout <- true
				return
			default:
				time.Sleep(1 * time.Second)
				data <- "生产数据"
			}
		}
	}()

	go func() {
		for {
			select {
			case <-isTimeout:
				log.Println("收到超时请求，关闭消费者协程")
				return
			case x := <-data:
				log.Println(x)

			}
		}
	}()

	time.Sleep(1 * time.Minute)
}

// 不能正常运行
func Test3() {
	isTimeout := make(chan bool)
	data := make(chan string)

	go func() {
		for {
			select {
			case <-isTimeout:
				log.Println("收到超时请求，关闭生产者协程")
				return
			default:
				time.Sleep(1 * time.Second)
				data <- "生产数据"
			}

		}
	}()

	// 设置消费者超时时间
	after := time.After(5 * time.Second)
	go func() {

		for {
			select {
			case x := <-data:
				log.Println("从data通道拿数据：", x)

			case <-after:
				log.Println("消费者超时，不再消费")
				isTimeout <- true
				return
			}
		}
	}()

	defer close(data)

	time.Sleep(60 * time.Second)
}

// 能正常运行
func Test4() {
	isTimeout := make(chan bool)
	data := make(chan string)

	go func() {
		for {
			select {
			case <-isTimeout:
				log.Println("收到超时请求，关闭生产者协程")
				return
			default:
				data <- "生产数据"
				time.Sleep(1 * time.Second)
			}

		}
	}()

	// 设置消费者超时时间
	after := time.After(5 * time.Second)
	go func() {

		for {
			select {
			case x := <-data:
				log.Println("从data通道拿数据：", x)

			case <-after:
				log.Println("消费者超时，不再消费")
				isTimeout <- true
				return
			}
		}
	}()

	defer close(data)

	time.Sleep(60 * time.Second)
}
