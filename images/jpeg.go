package images

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func CompressJpg(buf []byte, quality int) []byte {

	cropBuf, err := bimg.NewImage(buf).Process(bimg.Options{
		Quality: quality,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return cropBuf
}
