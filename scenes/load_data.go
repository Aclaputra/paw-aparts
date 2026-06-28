package scenes

import (
	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/aclaputra/paw-aparts/factory"
	"github.com/aclaputra/paw-aparts/layers"
	"github.com/aclaputra/paw-aparts/systems"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
)

func loadECSLevel1() *ecs.ECS {
	ecs := ecs.NewECS(donburi.NewWorld())

	return ecs
}

func loadECSMainMenu() *ecs.ECS {
	ecs := ecs.NewECS(donburi.NewWorld())

	ecs.AddSystem(systems.UpdateMenu)
	ecs.AddRenderer(layers.Default, systems.DrawMenu)

	factory.CreateMenu(ecs, components.MenuData{
		BackgroundImage: assets.GetEbitenImage(assets.Background_Png),
		Texts: []components.TextData{
			{
				Loc: &math.Vec2{X: float64((config.C.Width)/2) - 100, Y: float64(config.C.Height) - 50},
				Txt: "Press Any To Continue",
			},
		},
		Actions:            []int{components.PRESS_ANY_TO_CONTINUE},
		NextSceneTriggered: false,
	})

	return ecs
}
