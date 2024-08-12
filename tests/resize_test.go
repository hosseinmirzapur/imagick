package test

import (
	"image-processer/imagick"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResize(t *testing.T) {
	err := imagick.Resize("./assets/image.png", 1000, 1000)
	assert.Nil(t, err)
}

func TestForceResize(t *testing.T) {
	err := imagick.ForceResize("./assets/image.png", 1000, 1000)
	assert.Nil(t, err)
}
