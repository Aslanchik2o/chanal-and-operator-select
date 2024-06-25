package main

import (
	"fmt"
	"time"
)

func main() {
	//cnanal принимает string\int\bool и тд
	var msq chan int
	fmt.Println(msq)
	msq = make(chan int )
	//каналы работают как указатель наверное:\
	fmt.Println(msq)
go func() {
	time.Sleep(time.Second)
	msq <- 12
}()
value :=<- msq
fmt.Println(value)
	//это типа мы записываем в канал 
	// все это работает так же и с int хз насчет bool

	
}