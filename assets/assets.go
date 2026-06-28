package assets

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed background.png
	Background_Png []byte
	//go:embed fonts/excel.ttf
	ExcelFont []byte
)

// GetEbitenImage converts embedded []byte into *ebiten.Image
func GetEbitenImage(raw []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(raw))
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}
	return ebiten.NewImageFromImage(img)
}
