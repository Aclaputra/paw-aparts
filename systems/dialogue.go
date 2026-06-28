package systems

import (
	"image/color"

	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi/ecs"
)

func UpdateDialogue(ecs *ecs.ECS) {
	entry, ok := components.Dialogue.First(ecs.World)
	if !ok {
		return
	}
	d := components.Dialogue.Get(entry)

	if d.Active && !d.Done {
		// typing effect: reveal one character per frame
		if d.CurrentIndex < len(d.Text) {
			d.CurrentIndex++
		} else {
			d.Done = true
		}
	}
}

func DrawDialogue(ecs *ecs.ECS, screen *ebiten.Image) {
	entry, ok := components.Dialogue.First(ecs.World)
	if !ok {
		return
	}
	d := components.Dialogue.Get(entry)

	textOP := text.DrawOptions{}
	x := d.Loc.X
	y := d.Loc.Y
	textOP.GeoM.Translate(float64(x), float64(y))
	textOP.ColorScale.ScaleWithColor(color.White)
	textOP.LineSpacing = 12
	if d.Active {
		visibleText := d.Text[:d.CurrentIndex]
		text.Draw(screen, visibleText, &text.GoTextFace{
			Source: fonts.UIFaceSource,
			Size:   12,
		}, &textOP)
	}
}
