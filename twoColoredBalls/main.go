package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		res := twoColoredBalls()
		fmt.Printf(res)
	}
}

func twoColoredBalls() string {
	var balls [7]int
	//第一步，从1到33随机选择6个数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		balls[i] = r.Intn(33) + 1
	}
	//第二步，从1到16中随机选择一个数
	rb := rand.New(rand.NewSource(time.Now().UnixNano()))
	balls[6] = rb.Intn(16) + 1
	res := fmt.Sprintf("中奖号码为%v\n", balls)
	return res
}
