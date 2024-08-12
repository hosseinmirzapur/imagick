package imagick

import "github.com/h2non/bimg"

func Flip(path string) error {
	buff, err := readFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Flip()
	if err != nil {
		return err
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
