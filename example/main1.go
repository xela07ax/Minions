package main

import (
	"fmt"
	"goNewProject/Minions"
	"time"
)

type MD5Summer struct {
	filenameIn chan string
	done       chan struct{} //Сигнал, что пока закругляться
	stopX      chan bool     //Миньон завершила свою работу
	Wait       chan bool
}

func main() {
	fmt.Println("Hello and welcome")
	conveer := Minions.NewConveer("test1", CounterFunc)
	conveer.RunMinions(10)
	fmt.Println("Init ok ", conveer.GetCores())
	for i := 1; i <= 500; i++ {
		//fmt.Println("i =", 100/i)
		conveer.InputChan <- fmt.Sprintf("element%d.txt", i)
	}
	time.Sleep(1 * time.Millisecond)
	conveer.Kill()
	conveer.GetCores()
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("By and welcome, %v!\n", conveer.GetCores())
}

func CounterFunc(gopherNumber int, element interface{}, out chan interface{}) {
	fmt.Printf("\"gopher%d.txt\"i = %v\n", gopherNumber, element)
}

// Операция прервется с сообщениями