package main

import (
	"fmt"
	"time"
)

func main() {
	consumer := new(IntConsumer)
	zeroEvenOdd := new(ZeroEvenOdd)
	zeroEvenOdd.N = 50

	c1 := make(chan int) //代表打印0的通道
	c2 := make(chan int) //代表打印偶数的通道
	c3 := make(chan int) //代表打印奇数的通道

	//启动三个协程
	go func() {
		zeroEvenOdd.Zero(consumer, c1, c2, c3)
	}()
	go func() {
		zeroEvenOdd.Even(consumer, c1, c2, c3)
	}()
	go func() {
		zeroEvenOdd.Odd(consumer, c1, c2, c3)
	}()
	c1 <- 1 //启动打印0
	time.Sleep(10 * time.Second)
}

type IntConsumer struct {
}

func (c *IntConsumer) Accept(x int) {
	fmt.Print(x)
}

type ZeroEvenOdd struct {
	N int
}

func (self *ZeroEvenOdd) Zero(consumer *IntConsumer, c1, c2, c3 chan int) {
	for i := 0; i < self.N; i++ {
		select {
		case num := <-c1:
			consumer.Accept(0)
			if num%2 == 0 {
				c2 <- num
			} else {
				c3 <- num
			}
		}
	}
}

func (self *ZeroEvenOdd) Odd(consumer *IntConsumer, c1, c2, c3 chan int) {
	for i := 1; i <= self.N; i += 2 {
		select {
		case num := <-c3:
			consumer.Accept(num)
			num++
			c1 <- num
		}
	}
}

func (self *ZeroEvenOdd) Even(consumer *IntConsumer, c1, c2, c3 chan int) {
	for i := 2; i <= self.N; i += 2 {
		select {
		case num := <-c2:
			consumer.Accept(num)
			num++
			c1 <- num
		}
	}
}
