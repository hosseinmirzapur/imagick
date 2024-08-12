package imagick

import "github.com/h2non/bimg"

func Interpretate(path string, intrprt bimg.Interpretation) error {
	buff, err := readFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Colourspace(intrprt)
	if err != nil {
		return err
	}

	err = saveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
