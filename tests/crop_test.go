package test

import (
	"image-processer/imagick"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrop(t *testing.T) {
	err := imagick.Crop("./assets/image.png", 100, 100)
	assert.Nil(t, err)
}
