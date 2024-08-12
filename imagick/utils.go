package imagick

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/h2non/bimg"
)

func Extension(path string) string {
	return filepath.Ext(path)
}

func ReadFile(path string) ([]byte, error) {
	buff, err := bimg.Read(path)
	if err != nil {
		return nil, fmt.Errorf("unable to find the path specified")
	}

	return buff, nil

}

func SaveFile(image []byte, path string) error {
	savePath := fmt.Sprintf("%d%s", time.Now().UnixMicro(), Extension(path))
	err := bimg.Write(savePath, image)
	if err != nil {
		return fmt.Errorf("unable to write to specified path")
	}

	return nil
}
