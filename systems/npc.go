package systems

import (
	"image"

	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateNPC(ecs *ecs.ECS) {
	npcEntry, _ := components.NPC.First(ecs.World)
	playerEntry, _ := components.Player.First(ecs.World)
	npc := components.NPC.Get(npcEntry)
	npcObject := collision.GetObject(npcEntry)
	player := components.Player.Get(playerEntry)
	playerObject := collision.GetObject(playerEntry)

	if check := playerObject.Check(player.SpeedX, 0, "cat"); check != nil {
		dialogueEntry, _ := components.Dialogue.First(ecs.World)
		dialogue := components.Dialogue.Get(dialogueEntry)
		dialogue.Active = true
		dialogue.Text = "Press Y to Interact."
		dialogue.Loc.X, dialogue.Loc.Y = npcObject.Position.X, npcObject.Position.Y

		if inpututil.IsKeyJustPressed(ebiten.KeyY) ||
			ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton3) ||
			ebiten.IsGamepadButtonPressed(1, ebiten.GamepadButton3) {
			if dialogue.Done {
				dialogue.Done = false
			}
			npc.FacingRight = true
			dialogue.Text = ""
			dialogue.Active = false
			dialogue.CurrentIndex = 0

			dialogue.NextSceneTriggered = true
		}
	}

	npc.HandleAnimationCount()
}

func DrawNPC(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.NPC.Each(ecs.World, func(e *donburi.Entry) {
		npc := components.NPC.Get(e)
		o := collision.GetObject(e)

		op := ebiten.DrawImageOptions{}
		scale := 1.0
		var frameOX, frameOY int
		frameWidth, frameHeight :=
			npc.CurrentImage.Bounds().Dx()/int(npc.Animation.FrameTotal.X),
			npc.CurrentImage.Bounds().Dy()/int(npc.Animation.FrameTotal.Y)

		i := (npc.GetAnimationCount() / (int(npc.Animation.FrameTotal.X) * int(npc.Animation.FrameTotal.Y)) % npc.Animation.FrameCount)

		var sx, sy int
		switch npc.Animation.State {
		case "idle":
			sx = frameOX + i*frameWidth
			sy = frameOY
		default:
			sx = frameOX + i*frameWidth
			sy = frameOY
		}

		if npc.FacingRight {
			op.GeoM.Scale(-scale, scale)
			op.GeoM.Translate(
				o.Position.X+float64(frameWidth)*scale,
				o.Position.Y,
			)
		} else {
			op.GeoM.Scale(scale, scale)
			op.GeoM.Translate(
				o.Position.X,
				o.Position.Y,
			)
		}

		screen.DrawImage(
			npc.CurrentImage.SubImage(
				image.Rect(
					sx, sy, sx+frameWidth, sy+frameHeight,
				),
			).(*ebiten.Image), &op)
	})
}
