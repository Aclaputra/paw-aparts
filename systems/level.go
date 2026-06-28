package systems

import (
	"image/color"

	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi/ecs"
)

func DrawPlatform(ecs *ecs.ECS, screen *ebiten.Image) {
	for e := range tags.Platform.Iter(ecs.World) {
		o := collision.GetObject(e)
		drawColor := color.RGBA{180, 100, 0, 255}
		ebitenutil.DrawRect(screen, o.Position.X, o.Position.Y, o.Size.X, o.Size.Y, drawColor)
	}
}

func DrawWall(ecs *ecs.ECS, screen *ebiten.Image) {
	background := assets.GetEbitenImage(assets.BGGameDark_Png)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := background.Bounds().Dx(), background.Bounds().Dy()

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

	screen.DrawImage(background, op)

	for e := range tags.Wall.Iter(ecs.World) {
		o := collision.GetObject(e)
		drawColor := color.RGBA{20, 60, 20, 255}
		ebitenutil.DrawRect(screen, o.Position.X, o.Position.Y, o.Size.X, o.Size.Y, drawColor)
	}
}
