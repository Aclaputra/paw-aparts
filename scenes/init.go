package scenes

import (
	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/aclaputra/paw-aparts/factory"
	"github.com/aclaputra/paw-aparts/layers"
	"github.com/aclaputra/paw-aparts/systems"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
)

func loadECSLevel1() *ecs.ECS {
	ecs := ecs.NewECS(donburi.NewWorld())

	ecs.AddRenderer(layers.Default, systems.DrawPlatform)

	gw, gh := float64(config.C.Width), float64(config.C.Height)
	space := factory.CreateSpace(ecs)

	collision.Add(space,
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2), gh-128, 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2)-128, gh-128+64, 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2), gh-(64*6), 128, 16, "platform")),
		factory.CreatePlatform(ecs, resolv.NewObject((gw/2)-128, gh-(64*7), 128, 16, "platform")),
	)

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
				Loc: &math.Vec2{X: float64((config.C.Width)/2) - 150, Y: float64(config.C.Height) - 50},
				Txt: "Press Enter or Space To Continue",
			},
		},
		Actions:            []int{components.PRESS_ANY_TO_CONTINUE},
		NextSceneTriggered: false,
	})

	return ecs
}
