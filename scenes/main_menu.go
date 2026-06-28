package scenes

import (
	"image/color"
	_ "image/png"

	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/aclaputra/paw-aparts/factory"
	"github.com/aclaputra/paw-aparts/game"
	"github.com/aclaputra/paw-aparts/layers"
	"github.com/aclaputra/paw-aparts/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
)

func createECSMainMenu() *ecs.ECS {
	ecs := ecs.NewECS(donburi.NewWorld())

	ecs.AddSystem(systems.UpdateMusic)
	ecs.AddSystem(systems.UpdateMenu)
	ecs.AddRenderer(layers.Default, systems.DrawMenu)

	factory.InitMusic(ecs, assets.MusicMainMenu)
	factory.CreateMenu(ecs, components.MenuData{
		BackgroundImage: assets.GetEbitenImage(assets.Background_Png),
		Texts: []components.TextData{
			{
				Loc: &math.Vec2{X: float64((config.C.Width)/2) - 150, Y: float64(config.C.Height) - 50},
				Txt: "Press Enter or Space To Continue",
			},
		},
		Actions:            []int{components.PRESS_ANY_TO_CONTINUE},
		NextSceneTriggered: false,
	})

	return ecs
}

type MainMenuScene struct {
	ecs *ecs.ECS
}

func NewMainMenuScene() *MainMenuScene {
	return &MainMenuScene{
		ecs: createECSMainMenu(),
	}
}

func (m *MainMenuScene) Update(g *game.Game) {
	m.ecs.Update()

	if ent, ok := components.Menu.First(m.ecs.World); ok {
		menu := components.Menu.Get(ent)
		if menu.NextSceneTriggered {
			menu.NextSceneTriggered = false

			systems.DisposeMusic(m.ecs)
			g.ChangeScene(NewLevel1Scene())
		}
	}
}

func (m *MainMenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{60, 60, 100, 255})
	m.ecs.Draw(screen)
}
