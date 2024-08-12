package test

import (
	"image-processer/imagick"
	"testing"

	"github.com/h2non/bimg"
	"github.com/stretchr/testify/assert"
)

func TestColors(t *testing.T) {
	err := imagick.Interpret("./assets/image.png", bimg.InterpretationBW)
	assert.Nil(t, err)
}
