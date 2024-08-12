package test

import (
	"image-processer/imagick"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlip(t *testing.T) {
	err := imagick.Flip("./assets/image.png")
	assert.Nil(t, err)
}
