package main

import (
	"fmt"
	"io/ioutil"

	"github.com/vinaymavi/exampic.xyz/images"
)

func main() {
	fmt.Println("Hello, World!")
	buf, err := ioutil.ReadFile("./test/1millidollars-3JMMlYgTyo8-unsplash.jpg")

	if err != nil {
		fmt.Println(err)
	}
	cbuf := images.CompressJpg(buf, 10)

	err = ioutil.WriteFile("./test/1millidollars-3JMMlYgTyo8-unsplash-compressed.jpg", cbuf, 0644)

	if err != nil {
		fmt.Println(err)
	}
	images.CompressPng()
}
