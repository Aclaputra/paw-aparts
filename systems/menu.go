package systems

import (
	"image/color"
	"log"

	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi/ecs"
)

func UpdateMenu(ecs *ecs.ECS) {
	if _, ok := components.Menu.First(ecs.World); !ok {
		log.Fatal("UpdateMenu: cannot load menu")
	}

	ent, _ := components.Menu.First(ecs.World)
	menu := components.Menu.Get(ent)

	for _, action := range menu.Actions {
		switch action {
		case components.PRESS_ANY_TO_CONTINUE:
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
				inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
				inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
				log.Println("Any key pressed → continue")
				// next scene or close menu
				menu.NextSceneTriggered = true
			}
		case components.PRESS_ENTER_TO_CONTINUE:
			if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
				log.Println("Enter pressed → continue")
				// next scene or close menu
				menu.NextSceneTriggered = true
			}
		default:
			log.Println("UpdateMenu: menu actions not found")
		}
	}
}

func DrawMenu(ecs *ecs.ECS, screen *ebiten.Image) {
	if _, ok := components.Menu.First(ecs.World); !ok {
		log.Fatal("DrawMenu: cannot load menu")
	}

	ent, _ := components.Menu.First(ecs.World)
	menu := components.Menu.Get(ent)

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := menu.BackgroundImage.Bounds().Dx(), menu.BackgroundImage.Bounds().Dy()

	scaleX := float64(sw) / float64(bw)
	scaleY := float64(sh) / float64(bh)

	scale := scaleX
	if scaleY > scaleX {
		scale = scaleY
	}

	op.GeoM.Scale(scale, scale)

	x := float64(sw)/2 - float64(bw)*scale/2
	y := float64(sh)/2 - float64(bh)*scale/2
	op.GeoM.Translate(x, y)

	screen.DrawImage(menu.BackgroundImage, op)

	for _, tData := range menu.Texts {
		legacyFace := fonts.Excel.Get()

		v2Face := text.NewGoXFace(legacyFace)

		textOp := &text.DrawOptions{}
		textOp.GeoM.Translate(tData.Loc.X, tData.Loc.Y)
		textOp.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, tData.Txt, v2Face, textOp)
	}
}
