package imagick

import (
	"fmt"

	"github.com/h2non/bimg"
)

func ToFormat(path string, format bimg.ImageType) error {
	buff, err := bimg.Read(path)
	if err != nil {
		return fmt.Errorf("unable to find the file in the path provided\n")
	}

	newImage, err := bimg.NewImage(buff).Convert(format)
	if err != nil {
		return fmt.Errorf("unable to convert the image into the format provided\n")
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
