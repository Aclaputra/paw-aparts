package scenes

import (
	"image/color"
	"strconv"

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

func createECSResultMenuScene(prevLevelName string, timeTaken float64, nextLevelName string) *ecs.ECS {
	ecs := ecs.NewECS(donburi.NewWorld())

	ecs.AddSystem(systems.UpdateMusic)
	ecs.AddSystem(systems.UpdateMenu)
	ecs.AddRenderer(layers.Default, systems.DrawMenu)

	var texts []components.TextData
	textList := []components.TextData{
		{
			Loc: &math.Vec2{X: float64((config.C.Width)/2) - 125, Y: 100},
			Txt: "Congrats You Won the" + " - " + prevLevelName,
		},
		{
			Loc: &math.Vec2{X: float64((config.C.Width)/2) - 125, Y: 150},
			Txt: "Time Taken" + ": " + strconv.Itoa(int(timeTaken)) + " Sec.",
		},
		{
			Loc: &math.Vec2{X: float64((config.C.Width)/2) - 180, Y: float64(config.C.Height) - 50},
			Txt: "Press Enter or Space To Continue" + " - " + nextLevelName,
		},
	}

	textEnd := []components.TextData{
		{
			Loc: &math.Vec2{X: float64((config.C.Width)/2) - 125, Y: 100},
			Txt: "You Won! Congrats!!" + " - " + prevLevelName,
		},
		{
			Loc: &math.Vec2{X: float64((config.C.Width)/2) - 125, Y: 150},
			Txt: "Thanks for Playing PAW APARTS :P",
		},
		{
			Loc: &math.Vec2{X: float64((config.C.Width)/2) - 180, Y: float64(config.C.Height) - 50},
			Txt: "Press Enter or Space To Continue" + " - " + nextLevelName,
		},
	}

	if prevLevelName != "End" && nextLevelName != "End" {
		texts = textList
	} else {
		texts = textEnd
	}

	factory.InitMusic(ecs, assets.MusicMainMenu)
	factory.CreateMenu(ecs, components.MenuData{
		BackgroundImage:    assets.GetEbitenImage(assets.Background_Png),
		Texts:              texts,
		Actions:            []int{components.PRESS_ANY_TO_CONTINUE},
		NextSceneTriggered: false,
	})

	return ecs
}

type ResultMenuScene struct {
	ecs       *ecs.ECS
	nextScene string
}

func NewResultMenuScene(prevLevelName string, timeTaken float64, nextScene string) *ResultMenuScene {
	return &ResultMenuScene{
		ecs:       createECSResultMenuScene(prevLevelName, timeTaken, nextScene),
		nextScene: nextScene,
	}
}

func (r *ResultMenuScene) Update(g *game.Game) {
	r.ecs.Update()

	if ent, ok := components.Menu.First(r.ecs.World); ok {
		menu := components.Menu.Get(ent)
		if menu.NextSceneTriggered {
			menu.NextSceneTriggered = false

			systems.DisposeMusic(r.ecs)
			switch r.nextScene {
			case "Level 2":
				g.ChangeScene(NewLevel2Scene())
			case "Thanks":
				g.ChangeScene(NewResultMenuScene("End", 0, "End"))
			case "End":
				g.ChangeScene(NewMainMenuScene())
			}
		}
	}
}

func (r *ResultMenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{60, 60, 100, 255})
	r.ecs.Draw(screen)
}
