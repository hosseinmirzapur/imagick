package imagick

import "path/filepath"

func extension(path string) string {
	return filepath.Ext(path)
}
