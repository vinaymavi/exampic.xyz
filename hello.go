package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"strings"
	"syscall/js"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func jsonWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return js.ValueOf(map[string]interface{}{
			"foo": "bar",
		})
	})
}

func helloWorld() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Get("updateH3").Invoke("WASM Hello World", "Above message from WASM GO Lang")
		return nil
	})
}

func cropImage() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go cropImageGoRoutine(args[0].String())
		return nil
	})
}

func cropImageGoRoutine(imageByteStirng string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageByteStirng))
	img, err := jpeg.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	cropSize := image.Rect(0, 0, width/2+100, width/2+100)
	cropSize = cropSize.Add(image.Point{100, 80})
	croppedImage := img.(SubImager).SubImage(cropSize)
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, croppedImage, nil)
	cropedBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	js.Global().Get("createImageElement").Invoke(cropedBase64, imageByteStirng)
}

func copyFile() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fileName := args[0].String()
		fmt.Println("File name:", fileName)

		return nil
	})
}

func main() {
	fmt.Println("Hello, 世界")
	js.Global().Set("jsonWrapper", jsonWrapper())
	js.Global().Set("helloWorld", helloWorld())
	js.Global().Set("cropImage", cropImage())
	js.Global().Set("copyFile", copyFile())
	select {}
}
