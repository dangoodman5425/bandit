package managers

import (
	"io"
	"image"
	"log"
)

func ReadImage(r io.Reader)	image.Image {
	img, _, err := image.Decode(r)
	if err != nil {
		log.Panic(err)
	}
	return img
}