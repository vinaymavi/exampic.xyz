package main

import (
	"fmt"
	"syscall/js"

	"github.com/vinaymavi/exampic.xyz/images"
)

func compressJpeg() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var buf []byte = []byte(args[0].String())
		quality := args[1].Int()
		cropBuf := images.CompressJpg(buf, quality)
		return cropBuf
	})

}

func main() {
	fmt.Println("Hello, World!")
	js.Global().Set("compressJpeg", compressJpeg())
	select {}
}
