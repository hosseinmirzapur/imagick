package imagick

import (
	"github.com/h2non/bimg"
)

func Rotate(path string, rotate bimg.Angle) error {
	buffer, err := readFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buffer).Rotate(rotate)
	if err != nil {
		return err
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
