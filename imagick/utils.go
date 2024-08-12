package imagick

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/h2non/bimg"
)

func extension(path string) string {
	return filepath.Ext(path)
}

func readFile(path string) ([]byte, error) {
	buff, err := bimg.Read(path)
	if err != nil {
		return nil, fmt.Errorf("unable to find the path specified\n")
	}

	return buff, nil

}

func saveFile(image []byte, path string) error {
	savePath := fmt.Sprintf("%d.%s", time.Now().UnixMicro(), extension(path))
	err := bimg.Write(savePath, image)
	if err != nil {
		return fmt.Errorf("unable to write to specified path\n")
	}

	return nil
}
