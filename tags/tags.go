package tags

import "github.com/yohamta/donburi"

var (
	Player   = donburi.NewTag().SetName("player")
	Platform = donburi.NewTag().SetName("platform")
	Wall     = donburi.NewTag().SetName("wall")
	NPC      = donburi.NewTag().SetName("npc")
	Dialogue = donburi.NewTag().SetName("dialogue")
	Music    = donburi.NewTag().SetName("music")
)
