package main

import (
	"bytes"
	"log"
	"math/rand"
	"time"

	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/aclaputra/paw-aparts/fonts"
	"github.com/aclaputra/paw-aparts/game"
	"github.com/aclaputra/paw-aparts/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func init() {
	fonts.LoadFont(fonts.Excel, assets.ExcelFont)
	uiFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	fonts.SetUIFaceSource(uiFaceSource)
}

func main() {
	game := &game.Game{}

	game.ChangeScene(scenes.NewMainMenuScene())
	ebiten.SetWindowSize(config.C.Width, config.C.Height)
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowTitle(config.C.GameTitle)
	rand.Seed(time.Now().UTC().UnixNano())
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
