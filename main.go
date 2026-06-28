package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/aclaputra/paw-aparts/config"
	"github.com/aclaputra/paw-aparts/fonts"
	"github.com/aclaputra/paw-aparts/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	//go:embed assets/fonts/excel.ttf
	excelFont []byte
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	bounds image.Rectangle
	scene  Scene
}

func NewGame() *Game {
	fonts.LoadFont(fonts.Excel, excelFont)
	uiFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	fonts.SetUIFaceSource(uiFaceSource)

	return &Game{
		bounds: image.Rectangle{},
		scene:  &scenes.MainMenuScene{},
	}
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func main() {
	ebiten.SetWindowSize(config.C.Width, config.C.Height)
	ebiten.SetWindowResizable(false)
	rand.Seed(time.Now().UTC().UnixNano())
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
