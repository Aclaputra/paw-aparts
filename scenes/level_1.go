package scenes

import (
	"image/color"
	"time"

	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/components"
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

	ecs.AddSystem(systems.UpdateMusic)
	ecs.AddSystem(systems.UpdateDialogue)
	ecs.AddSystem(systems.UpdateNPC)
	ecs.AddSystem(systems.UpdatePlayer)

	ecs.AddRenderer(layers.Default, systems.DrawWall)
	ecs.AddRenderer(layers.Default, systems.DrawDialogue)
	ecs.AddRenderer(layers.Default, systems.DrawNPC)
	ecs.AddRenderer(layers.Default, systems.DrawPlatform)
	ecs.AddRenderer(layers.Default, systems.DrawPlayer)

	gw, gh := float64(config.C.Width), float64(config.C.Height)
	space := factory.CreateSpace(ecs)

	factory.InitMusic(ecs, assets.MusicGame)
	factory.CreateDialogue(ecs)
	collision.Add(space,
		factory.CreateNPC(ecs, &math.Vec2{
			X: (float64(config.C.Width) / 2) - (80 + 64),
			Y: float64(config.C.Height) - (128 * 3) + 16,
		}, "cat"),
		factory.CreatePlayer(ecs, &math.Vec2{X: gw / 2, Y: gh / 2}),
		factory.CreateWall(ecs, resolv.NewObject((gw/2)-128, gh-(256*1.2), 32, 130*2, "solid")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2), gh-128, 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2)-128, gh-128+64, 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2)-128, gh-(64*5), 128, 16, "platform")),
	)

	return ecs
}

type Level1Scene struct {
	ecs        *ecs.ECS
	lastUpdate time.Time
	elapsed    float64
}

func NewLevel1Scene() *Level1Scene {
	return &Level1Scene{
		ecs: createECSLevel1(),
	}
}

func (l1 *Level1Scene) Update(g *game.Game) {
	if l1.lastUpdate.IsZero() {
		l1.lastUpdate = time.Now()
	}

	now := time.Now()
	dt := now.Sub(l1.lastUpdate).Seconds()
	l1.lastUpdate = now

	l1.elapsed += dt

	l1.ecs.Update()

	if ent, ok := components.Dialogue.First(l1.ecs.World); ok {
		menu := components.Dialogue.Get(ent)
		if menu.NextSceneTriggered {
			menu.NextSceneTriggered = false

			systems.DisposeMusic(l1.ecs)
			g.ChangeScene(NewResultMenuScene("Level 1", l1.elapsed, "Level 2"))
		}
	}
}

func (l1 *Level1Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{60, 60, 100, 255})
	l1.ecs.Draw(screen)
}
