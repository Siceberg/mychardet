package main

import (
	"log"

	"gitlab.com/siceberg/chardet"
)

func main() {
	log.Println(chardet.Mostlike([]byte("test")))
}
