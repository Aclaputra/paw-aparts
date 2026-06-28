package systems

import (
	"image/color"

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
	for e := range tags.Wall.Iter(ecs.World) {
		o := collision.GetObject(e)
		drawColor := color.RGBA{20, 60, 20, 255}
		ebitenutil.DrawRect(screen, o.Position.X, o.Position.Y, o.Size.X, o.Size.Y, drawColor)
	}
}
