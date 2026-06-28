package scenes

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type MainMenuScene struct {
	ecs  *ecs.ECS
	once sync.Once
}

func (m *MainMenuScene) Update() {
	m.once.Do(m.configure)
	m.ecs.Update()
}

func (m *MainMenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{60, 60, 100, 255})
	m.ecs.Draw(screen)
}

func (m *MainMenuScene) configure() {
	ecs := ecs.NewECS(donburi.NewWorld())

	m.ecs = ecs
}
