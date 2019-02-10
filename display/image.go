package display

import (
	"image"
	"image/color"
	"os"
	"tfl-chromecast/tfl"
)

var ()

// ffmpeg -f v4l2 -s 320x240 -r 25 -i /dev/video0 -f alsa -ac 1 -i hw:0 http://localhost:8090/feed1.ffm
func GenerateImage(stops []tfl.BusStop) (image.Image, error) {

	img := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
	img.Set(2, 3, color.RGBA{255, 0, 0, 255})

}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	return img, err
}
