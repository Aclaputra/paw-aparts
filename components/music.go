package components

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/yohamta/donburi"
)

type MusicData struct {
	Background *audio.Player
}

var Music = donburi.NewComponentType[MusicData]()
