package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main(){
	fmt.Println("samaller text message")
	message := []byte(" I live in Bialystok and I love having fun")
	
	fmt.Println("MDS: %x", md5.Sum(message))
	fmt.Println("SHA256: %x", sha256.Sum256(message))
	fmt.Println("SHA512: %x", sha512.Sum512(message))

	fmt.Println("larger text message")
	message = []byte("I live in Bialystok and I love having fun and here are a few more messages and texts and font etc. etc.")

	fmt.Println("MDS: %x", md5.Sum(message))
	fmt.Println("SHA256: %x", sha256.Sum256(message))
	fmt.Println("SHA512: %x", sha512.Sum512(message))

}