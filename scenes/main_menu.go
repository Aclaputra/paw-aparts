package scenes

import (
	"image/color"
	_ "image/png"

	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

type MainMenuScene struct {
	ecs *ecs.ECS
}

func NewMainMenuScene() *MainMenuScene {
	return &MainMenuScene{
		ecs: loadECSMainMenu(),
	}
}

func (m *MainMenuScene) Update(g *game.Game) {
	m.ecs.Update()

	if ent, ok := components.Menu.First(m.ecs.World); ok {
		menu := components.Menu.Get(ent)
		if menu.NextSceneTriggered {
			menu.NextSceneTriggered = false

			g.ChangeScene(NewLevel1Scene())
		}
	}
}

func (m *MainMenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{60, 60, 100, 255})
	m.ecs.Draw(screen)
}
