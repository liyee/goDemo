package main

import "fmt"

type RoundServer struct {
	currentIndex int
	serverList   []string
}

func (this *RoundServer) Get() {
	server := this.serverList[this.currentIndex]
	this.currentIndex = (this.currentIndex + 1) % len(this.serverList)
	fmt.Println(server)
}

func main() {
	roud := &RoundServer{0, []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}}

	for i := 0; i < 10; i++ {
		roud.Get()
	}
}
