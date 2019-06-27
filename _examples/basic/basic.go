package main

import (
	"fmt"

	jptorrent "github.com/ken8203/jptorrent"
)

func main() {
	err := jptorrent.Download("5SGT4VGE1B")
	if err != nil {
		fmt.Println(err)
	}
}
