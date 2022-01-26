package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func PutIn(ch chan string){
	var num =1
	for i:=0;i<4;i++{
		ch <- Arr[i]
		if i ==3 {
			i = -1
			num++
		}
		if num == 11{
			break
		}
	}
}
func Work(ch chan string){                //用Work函数读取channel的内容
	for i:=0;i<40;i++ {
		a := <-ch
		fmt.Println(a)
	}
	wg.Done()
}
var Arr = map[int]string{
          0:"张三",
          1:"李四",
          2:"王五",
          3:"赵六",
}
func main() {
	var ch = make(chan string,1)
	wg.Add(1)
	go PutIn(ch)
	go Work(ch)
	wg.Wait()
}
