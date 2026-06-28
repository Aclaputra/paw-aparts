package factory

import (
	"github.com/aclaputra/paw-aparts/archetypes"
	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
)

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
