package imagick

import "github.com/h2non/bimg"

func Flip(path string) error {
	buff, err := ReadFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Flip()
	if err != nil {
		return err
	}

	err = SaveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
