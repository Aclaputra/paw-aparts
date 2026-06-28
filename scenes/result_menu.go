package scenes

import (
	"image/color"

	"github.com/aclaputra/paw-aparts/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func createECSResultMenuScene() *ecs.ECS {
	ecs := ecs.NewECS(donburi.NewWorld())

	return ecs
}

type ResultMenuScene struct {
	ecs *ecs.ECS
}

func NewResultMenuScene() *ResultMenuScene {
	return &ResultMenuScene{
		ecs: createECSResultMenuScene(),
	}
}

func (r *ResultMenuScene) Update(g *game.Game) {
	r.ecs.Update()
}

func (r *ResultMenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{60, 60, 100, 255})
	r.ecs.Draw(screen)
}
