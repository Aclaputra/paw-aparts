package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Bounds image.Rectangle
	Scene  Scene
}

type Scene interface {
	Update(g *Game)
	Draw(screen *ebiten.Image)
}

func (g *Game) ChangeScene(newScene Scene) {
	g.Scene = newScene
}

func (g *Game) Update() error {
	g.Scene.Update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.Scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.Bounds = image.Rect(0, 0, width, height)
	return width, height
}
