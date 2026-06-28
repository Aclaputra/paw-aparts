package factory

import (
	"github.com/aclaputra/paw-aparts/archetypes"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateMenu(ecs *ecs.ECS, data components.MenuData) *donburi.Entry {
	menu := archetypes.Menu.Spawn(ecs)

	components.Menu.SetValue(menu, data)

	return menu
}
