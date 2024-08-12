package test

import (
	"image-processer/imagick"
	"testing"

	"github.com/h2non/bimg"
	"github.com/stretchr/testify/assert"
)

func TestWatermark(t *testing.T) {
	wmConfig := bimg.Watermark{
		Width:      400,
		Text:       "Hossein Mirzapur",
		Opacity:    0.45,
		DPI:        100,
		Margin:     150,
		Font:       "sans bold 14",
		Background: bimg.ColorBlack,
	}
	err := imagick.Watermark("./assets/image.png", wmConfig)
	assert.Nil(t, err)
}
