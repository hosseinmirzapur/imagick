package imagick

import (
	"github.com/h2non/bimg"
)

func ToFormat(path string, format bimg.ImageType) error {
	buff, err := readFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Convert(format)
	if err != nil {
		return err
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
