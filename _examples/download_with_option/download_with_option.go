package main

import (
	"fmt"

	jptorrent "github.com/ken8203/jptorrent"
	options "github.com/ken8203/jptorrent/options"
)

func main() {
	opt := options.Option{
		Location: "./so-good.torrent",
	}
	err := jptorrent.Download("5SGT4VGE1B", opt)
	if err != nil {
		fmt.Println(err)
	}
}
