package scenes

import (
	"image/color"

	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/aclaputra/paw-aparts/factory"
	"github.com/aclaputra/paw-aparts/game"
	"github.com/aclaputra/paw-aparts/layers"
	"github.com/aclaputra/paw-aparts/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
)

func createECSLevel1() *ecs.ECS {
	ecs := ecs.NewECS(donburi.NewWorld())

	ecs.AddSystem(systems.UpdatePlayer)
	
	ecs.AddRenderer(layers.Default, systems.DrawPlayer)
	ecs.AddRenderer(layers.Default, systems.DrawPlatform)

	gw, gh := float64(config.C.Width), float64(config.C.Height)
	space := factory.CreateSpace(ecs)

	collision.Add(space,
		factory.CreatePlayer(ecs, &math.Vec2{X: gw / 2, Y: gh / 2}),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2), gh-128, 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2)-128, gh-128+64, 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2), gh-(64*6), 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2)-128, gh-(64*7), 128, 16, "platform")),
	)

	return ecs
}

type Level1Scene struct {
	ecs *ecs.ECS
}

func NewLevel1Scene() *Level1Scene {
	return &Level1Scene{
		ecs: createECSLevel1(),
	}
}

func (l1 *Level1Scene) Update(g *game.Game) {
	l1.ecs.Update()
}

func (l1 *Level1Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{60, 60, 100, 255})
	l1.ecs.Draw(screen)
}
