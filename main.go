package main

import (
	"fmt"
)

func main() {
	msq := make(chan string, 4)//буферезованный канал
	msq <- "Ниняно"
	msq <- "Ниняно 2"
fmt.Println(<-msq)
fmt.Println(<-msq)
	
}