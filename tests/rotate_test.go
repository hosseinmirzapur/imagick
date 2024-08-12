package test

import (
	"image-processer/imagick"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	err := imagick.Rotate("./assets/image.png", 180)
	assert.Nil(t, err)
}
