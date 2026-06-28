package systems

import (
	"log"

	"github.com/aclaputra/paw-aparts/components"
	"github.com/yohamta/donburi/ecs"
)

func UpdateMusic(ecs *ecs.ECS) {
	world := ecs.World
	if entry, ok := components.Music.First(world); ok {
		musicData := components.Music.Get(entry)

		musicData.Background.Play()
	} else {
		log.Println("UpdateMusic: cannot update music")
	}
}

func DisposeMusic(ecs *ecs.ECS) {
	world := ecs.World
	if entry, ok := components.Music.First(world); ok {
		musicData := components.Music.Get(entry)

		if musicData.Background != nil {
			musicData.Background.Pause()
			musicData.Background.Close() // This releases Ebitengine's audio stream memory
		}

		// Remove the entity from the Donburi world so it doesn't linger
		world.Remove(entry.Entity())
	}
}
