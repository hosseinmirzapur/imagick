package imagick

import (
	"fmt"
	"time"

	"github.com/h2non/bimg"
)

func Rotate(path string, rotate bimg.Angle) error {
	buffer, err := readFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buffer).Rotate(rotate)
	if err != nil {
		return fmt.Errorf("cannot rotate the image\n")
	}

	savePath := fmt.Sprintf("%d.%s", time.Now().UnixMicro(), extension(path))
	err = bimg.Write(savePath, newImage)
	if err != nil {
		return fmt.Errorf("cannot save the result image in specified path\n")
	}

	return nil
}
