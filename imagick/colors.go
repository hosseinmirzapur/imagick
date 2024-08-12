package imagick

import "github.com/h2non/bimg"

func Interpret(path string, intrprt bimg.Interpretation) error {
	buff, err := ReadFile(path)
	if err != nil {
		return err
	}

	newImage, err := bimg.NewImage(buff).Colourspace(intrprt)
	if err != nil {
		return err
	}

	err = SaveFile(newImage, path)
	if err != nil {
		return err
	}

	return nil
}
