package imagick

import "github.com/h2non/bimg"

func Watermark(path string, wmConf bimg.Watermark) error {
	buff, err := ReadFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Watermark(wmConf)
	if err != nil {
		return err
	}

	err = SaveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
