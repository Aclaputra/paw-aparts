package factory

import (
	"bytes"
	"log"

	"github.com/aclaputra/paw-aparts/components"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/yohamta/donburi/ecs"
)

const sampleRate = 44100

func InitMusic(ecs *ecs.ECS, oggBytes []byte) {
	world := ecs.World

	ctx := audio.CurrentContext()
	if ctx == nil {
		ctx = audio.NewContext(sampleRate)
	}

	// Decode and create the player normally
	stream, err := vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(oggBytes))
	if err != nil {
		log.Fatalf("failed to decode ogg: %v", err)
	}

	loopStream := audio.NewInfiniteLoop(stream, stream.Length())
	player, err := ctx.NewPlayer(loopStream)
	if err != nil {
		log.Fatalf("failed to create audio player: %v", err)
	}

	// Re-save Donburi world
	entry := world.Entry(world.Create(components.Music))
	components.Music.SetValue(entry, components.MusicData{
		Background: player,
	})

	player.Play()
}
