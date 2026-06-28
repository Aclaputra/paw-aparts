package scenes

import (
	"fmt"
	"image/color"

	"github.com/aclaputra/paw-aparts/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

type Level1Scene struct {
	ecs *ecs.ECS
}

func NewLevel1Scene() *Level1Scene {
	return &Level1Scene{
		ecs: loadECSLevel1(),
	}
}

func (l1 *Level1Scene) Update(g *game.Game) {
	fmt.Println("Youre on level 1")
	l1.ecs.Update()
}

func (l1 *Level1Scene) Draw(screen *ebiten.Image) {
	fmt.Println("Youre on level 1")
	screen.Fill(color.RGBA{60, 60, 100, 255})
	l1.ecs.Draw(screen)
}
