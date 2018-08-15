package thumbnail

import (
	"bytes"
	"image"
	"github.com/disintegration/imaging"
	"io"
	"log"
)

func CreateThumbnail(buffer []byte, maxWidth int, maxHeight int) (img image.Image, format string, err error) {
	reader := bytes.NewBuffer(buffer)
	srcImage, format, err := loadImageFromBuffer(reader)
	if err != nil {
		return nil, "", err
	}
	rectangle := srcImage.Bounds()

	if theImageRatioIsSmallerThanTheThumbnalBoundaries(rectangle, maxHeight, maxWidth) {
		img = imaging.Resize(srcImage, 0, maxHeight, imaging.Lanczos)
	} else {
		img = imaging.Resize(srcImage, maxWidth, 0, imaging.Lanczos)
	}
	return
}

func theImageRatioIsSmallerThanTheThumbnalBoundaries(rectangle image.Rectangle, maxHeight int, maxWidth int) bool {
	return rectangle.Dx()*maxHeight <= maxWidth*rectangle.Dy()
}

func loadImageFromBuffer(reader io.Reader) (img image.Image, format string, err error) {
	img, format, err = image.Decode(reader)
	if err != nil {
		log.Printf("Error while decoding file : %s", err)
		return nil, "", err
	}
	return
}
