package main

import (
	"fmt"

	"github.com/vinaymavi/exampic.xyz/images"
)

func main() {
	fmt.Println("Hello, World!")
	images.CompressJpg()
	images.CompressPng()
}
