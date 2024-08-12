package imagick

import "github.com/h2non/bimg"

func Watermark(path string, wmConf bimg.Watermark) error {
	buff, err := readFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Watermark(wmConf)
	if err != nil {
		return err
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
