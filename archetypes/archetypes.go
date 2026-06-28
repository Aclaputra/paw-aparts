package archetypes

import (
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/layers"
	"github.com/aclaputra/paw-aparts/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var (
	Menu = newArchetype(
		components.Menu,
	)
	Platform = newArchetype(
		tags.Platform,
		components.Object,
	)
	Player = newArchetype(
		tags.Player,
		components.Player,
		components.Object,
	)
	NPC = newArchetype(
		tags.NPC,
		components.NPC,
		components.Object,
	)
	Space = newArchetype(
		components.Space,
	)
	Wall = newArchetype(
		tags.Wall,
		components.Object,
	)
	Dialogue = newArchetype(
		tags.Dialogue,
		components.Dialogue,
	)
)

type archetype struct {
	components []donburi.IComponentType
}

func newArchetype(cs ...donburi.IComponentType) *archetype {
	return &archetype{
		components: cs,
	}
}

func (a *archetype) Spawn(ecs *ecs.ECS, cs ...donburi.IComponentType) *donburi.Entry {
	e := ecs.World.Entry(ecs.Create(
		layers.Default,
		append(a.components, cs...)...,
	))
	return e
}
