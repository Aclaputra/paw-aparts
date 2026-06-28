package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type MenuData struct {
	BackgroundImage    *ebiten.Image
	Texts              []TextData
	Actions            []int
	NextSceneTriggered bool
}

type TextData struct {
	Loc *math.Vec2
	Txt string
}

const (
	PRESS_ENTER_TO_CONTINUE int = iota
	PRESS_ANY_TO_CONTINUE
)

var Menu = donburi.NewComponentType[MenuData]()
