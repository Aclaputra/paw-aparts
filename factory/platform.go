package factory

import (
	"github.com/aclaputra/paw-aparts/archetypes"
	"github.com/aclaputra/paw-aparts/collision"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlatform(ecs *ecs.ECS, object *resolv.Object) *donburi.Entry {
	platform := archetypes.Platform.Spawn(ecs)
	collision.SetObject(platform, object)

	return platform
}
