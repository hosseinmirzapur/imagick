package test

import (
	"image-processer/imagick"
	"testing"

	"github.com/h2non/bimg"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	err := imagick.ToFormat("./assets/image.png", bimg.JPEG)
	assert.Nil(t, err)
}
