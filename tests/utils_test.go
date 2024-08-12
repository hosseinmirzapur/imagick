package test

import (
	"crypto/rand"
	"image-processer/imagick"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtension(t *testing.T) {
	assert.Equal(t, ".png", imagick.Extension("./assets/image.png"))
}

func TestReadFile(t *testing.T) {
	buff, err := imagick.ReadFile("./assets/image.png")
	assert.NotNil(t, buff)
	assert.Nil(t, err)
}

func TestSaveFile(t *testing.T) {

	bytes := make([]byte, 64_000)
	_, err := rand.Read(bytes)
	assert.Nil(t, err)

	err = imagick.SaveFile(bytes, "./assets/test.txt")
	assert.Nil(t, err)

}
