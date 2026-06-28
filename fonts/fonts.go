package fonts

import (
	"fmt"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

type FontName string

const (
	Excel FontName = "excel"
)

var UIFaceSource *text.GoTextFaceSource

func (f FontName) Get() font.Face {
	return getFont(f)
}

var (
	fonts = map[FontName]font.Face{}
)

func LoadFont(name FontName, ttf []byte) {
	fontData, _ := truetype.Parse(ttf)
	fonts[name] = truetype.NewFace(fontData, &truetype.Options{Size: 20})
}

func getFont(name FontName) font.Face {
	f, ok := fonts[name]
	if !ok {
		panic(fmt.Sprintf("Font %s not found", name))
	}
	return f
}

func SetUIFaceSource(ui *text.GoTextFaceSource) {
	UIFaceSource = ui
}
