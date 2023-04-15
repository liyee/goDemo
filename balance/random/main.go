package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type RandomBalance struct {
	curIndex   int
	serverList []string
}

func (this *RandomBalance) Add(server ...string) {
	this.serverList = append(this.serverList, server...)
}

func (this *RandomBalance) Get() string {
	index := len(this.serverList)
	curIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(index)))
	indexInt, _ := strconv.Atoi(curIndex.String())
	return this.serverList[indexInt]
}

func main() {
	balance := &RandomBalance{0, []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}}

	fmt.Println(balance.Get())
}
