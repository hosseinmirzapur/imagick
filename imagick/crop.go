package imagick

import "github.com/h2non/bimg"

func Crop(path string, width, height int) error {
	buff, err := readFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Crop(width, height, bimg.GravitySmart)
	if err != nil {
		return err
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
