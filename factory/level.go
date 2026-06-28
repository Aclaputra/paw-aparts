package factory

import (
	"bytes"
	"image"
	"log"

	"github.com/aclaputra/paw-aparts/archetypes"
	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
)

func CreateDialogue(ecs *ecs.ECS) *donburi.Entry {
	dialogue := archetypes.Dialogue.Spawn(ecs)

	components.Dialogue.SetValue(dialogue, components.DialogueData{})

	return dialogue
}

func CreateNPC(ecs *ecs.ECS, spawnPoint *math.Vec2, tags ...string) *donburi.Entry {
	npc := archetypes.NPC.Spawn(ecs)

	idleCatPng, _, err := image.Decode(bytes.NewReader(assets.IdleCat_Png))
	if err != nil {
		log.Println(err)
	}
	currentImage := ebiten.NewImageFromImage(idleCatPng)

	obj := resolv.NewObject(spawnPoint.X, spawnPoint.Y, 80, 64, tags...)
	collision.SetObject(npc, obj)
	// aniTotalX, aniTotalY, npc

	// should have just set the value in the factory for the animations
	components.NPC.SetValue(npc, components.NPCData{
		CurrentImage: currentImage,
		Animation: &components.Animation{
			State:      "idle",
			FrameCount: 8,
			FrameTotal: &math.Vec2{
				X: 8,
				Y: 1,
			},
		},
	})

	obj.SetShape(resolv.NewRectangle(0, 0, 80, 64))

	return npc
}

func CreatePlayer(ecs *ecs.ECS, spawnPoint *math.Vec2) *donburi.Entry {
	player := archetypes.Player.Spawn(ecs)

	// the 64 supposed to be not magic number
	obj := resolv.NewObject(spawnPoint.X, spawnPoint.Y, 80, 64)
	collision.SetObject(player, obj)
	// should have just set the value in the factory for the animations
	components.Player.SetValue(player, components.PlayerData{
		FacingRight: true,
		Animation: &components.Animation{
			FrameCount: 8,
			State:      "idle",
		}, // Idle -- should not magic number
	})

	// this too should not magic number
	obj.SetShape(resolv.NewRectangle(0, 0, 80, 64))

	return player
}

func CreateWall(ecs *ecs.ECS, obj *resolv.Object) *donburi.Entry {
	wall := archetypes.Wall.Spawn(ecs)
	collision.SetObject(wall, obj)
	return wall
}

func CreatePlatform(ecs *ecs.ECS, object *resolv.Object) *donburi.Entry {
	platform := archetypes.Platform.Spawn(ecs)
	collision.SetObject(platform, object)

	return platform
}

func CreateSpace(ecs *ecs.ECS) *donburi.Entry {
	space := archetypes.Space.Spawn(ecs)

	cfg := config.C
	spaceData := resolv.NewSpace(cfg.Width, cfg.Height, 32, 32)
	components.Space.Set(space, spaceData)

	return space
}
