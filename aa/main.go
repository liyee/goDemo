package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	fmt.Printf("%X", md5.Sum([]byte("fasdfadf")))
}
