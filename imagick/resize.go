package imagick

import (
	"fmt"

	"github.com/h2non/bimg"
)

func Resize(path string, width, height int) error {
	buffer, err := bimg.Read(path)
	if err != nil {
		return fmt.Errorf("unable to find the path specified\n")
	}

	newImage, err := bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		return fmt.Errorf("unable to resize image\n")
	}

	size, err := bimg.NewImage(newImage).Size()
	if size.Width != width || size.Height != height {
		return fmt.Errorf("unable to resize image properly\n")
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}

func ForceResize(path string, width, height int) error {
	buffer, err := bimg.Read(path)
	if err != nil {
		return fmt.Errorf("unable to find the path specified")
	}

	newImage, err := bimg.NewImage(buffer).ForceResize(width, height)
	if err != nil {
		return fmt.Errorf("unable to resize image")
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
