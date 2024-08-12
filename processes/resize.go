package processes

import (
	"fmt"
	"log"
	"time"

	"github.com/h2non/bimg"
)

func resize(path string, width, height int) {
	buffer, err := bimg.Read(path)
	if err != nil {
		log.Fatalln("unable to find the path specified")
	}

	newImage, err := bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		log.Fatalln("unable to resize image")
	}

	size, err := bimg.NewImage(newImage).Size()
	if size.Width != width || size.Height != height {
		log.Fatalln("unable to resize and scale image properly")
	}

	savePath := fmt.Sprintf("%d.%s", time.Now().UnixMicro(), extension(path))
	err = bimg.Write(savePath, newImage)
	if err != nil {
		log.Fatalln("unable to write to specified path")
	}
}

func forceResize(path string, width, height int) {
	buffer, err := bimg.Read(path)
	if err != nil {
		log.Fatalln("unable to find the path specified")
	}

	newImage, err := bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		log.Fatalln("unable to resize image")
	}

	savePath := fmt.Sprintf("%d.%s", time.Now().UnixMicro(), extension(path))
	err = bimg.Write(savePath, newImage)
	if err != nil {
		log.Fatalln("unable to write to specified path")
	}
}
