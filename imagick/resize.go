package imagick

import (
	"github.com/h2non/bimg"
)

func Resize(path string, width, height int) error {
	buff, err := ReadFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Resize(width, height)
	if err != nil {
		return err
	}

	size, err := bimg.NewImage(newImage).Size()
	if size.Width != width || size.Height != height {
		return err
	}

	err = SaveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}

func ForceResize(path string, width, height int) error {
	buff, err := ReadFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).ForceResize(width, height)
	if err != nil {
		return err
	}

	err = SaveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
