package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

var Space = donburi.NewComponentType[resolv.Space]()
var Object = donburi.NewComponentType[resolv.Object]()
var Player = donburi.NewComponentType[PlayerData]()
var NPC = donburi.NewComponentType[NPCData]()
var Dialogue = donburi.NewComponentType[DialogueData]()

type PlayerData struct {
	SpeedX          float64
	SpeedY          float64
	OnGround        *resolv.Object
	WallSliding     *resolv.Object
	FacingRight     bool
	IgnorePlatform  *resolv.Object
	Animation       *Animation
	RespawnCooldown int
}

type NPCData struct {
	Animation    *Animation
	CurrentImage *ebiten.Image
	FacingRight  bool
}

type DialogueData struct {
	Active       bool   // whether dialogue box is visible
	Text         string // full text to show
	CurrentIndex int    // typing effect index
	Done         bool   // finished typing
	Loc          math.Vec2
}

type Animation struct {
	State      string
	count      int
	FrameCount int
	FrameTotal *math.Vec2
}

func (n *NPCData) HandleAnimationCount() {
	n.Animation.count++
	if n.Animation.count%1000 == 0 {
		// reset
		n.Animation.count = 0
	}
}

func (n *NPCData) GetAnimationCount() int {
	return n.Animation.count
}

func (p *PlayerData) HandleAnimationCount() {
	p.Animation.count++
	if p.Animation.count%1000 == 0 {
		// reset
		p.Animation.count = 0
	}
}

func (p *PlayerData) GetAnimationCount() int {
	return p.Animation.count
}
