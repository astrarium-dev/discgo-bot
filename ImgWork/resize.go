package imgwork

import (
	"bytes"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"time"

	logger "github.com/hashbat-dev/discgo-bot/Logger"
	"github.com/nfnt/resize"
)

// Resize image takes an image.Image object, and a width and height, and passes modifies the passed in ResizeImage
// to give us the intended dimensions. Image is then encoded as .png format and returned as a single-use io.Reader.
// Note: This operation is destructive to image aspect ratios, so should only be used on things we do not mind being distorted.
func ResizeImage(guildId string, img image.Image, width uint) (io.Reader, error) {
	// Resize the image using the resize package
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()
	newHeight := uint(float64(originalHeight) * (float64(width) / float64(originalWidth)))
	resizedImg := resize.Resize(width, newHeight, img, resize.Bilinear)
	// Create a bytes buffer to write the PNG image to
	var buf bytes.Buffer
	err := png.Encode(&buf, resizedImg)
	if err != nil {
		logger.Error(guildId, err)
		return nil, err
	}

	logger.Debug(guildId, "Resized static image to %dx%d", newHeight, width)
	return &buf, nil
}

func ResizeGif(guildId string, gifImg *gif.GIF, width uint, height uint) (io.Reader, error) {
	// Resize each frame
	startTime := time.Now()
	for i, frame := range gifImg.Image {
		resizedFrame := resize.Resize(width, height, frame, resize.Lanczos3)

		// Convert the resized frame to *image.Paletted
		palettedFrame := image.NewPaletted(resizedFrame.Bounds(), gifImg.Image[i].Palette)

		// Draw the resized frame onto the paletted frame
		draw.FloydSteinberg.Draw(palettedFrame, palettedFrame.Rect, resizedFrame, image.Point{})

		// Update the frame in the GIF
		gifImg.Image[i] = palettedFrame
	}

	// Update the GIF configuration
	gifImg.Config.Width = int(width)
	gifImg.Config.Height = int(height)

	// Encode the resized frames back to GIF format
	var buf bytes.Buffer
	err := gif.EncodeAll(&buf, gifImg)
	if err != nil {
		logger.Error(guildId, err)
		return nil, err
	}

	logger.Info(guildId, "Resized GIF to %dx%d after %v", height, width, time.Since(startTime))
	return &buf, nil
}

func StretchImage(guildId string, img image.Image, width uint) (io.Reader, error) {
	bounds := img.Bounds()
	originalHeight := uint(bounds.Dy())
	newHeight := uint(float64(originalHeight) * 0.6)
	newWidth := uint(float64(width) * 1.3)

	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Bilinear)
	// Create a bytes buffer to write the PNG image to
	var buf bytes.Buffer
	err := png.Encode(&buf, resizedImg)
	if err != nil {
		logger.Error(guildId, err)
		return nil, err
	}

	logger.Debug(guildId, "Resized static image to %dx%d", originalHeight, newWidth)
	return &buf, nil
}
