package systems

import (
	"image"
	"math"

	"github.com/aclaputra/paw-aparts/assets"
	"github.com/aclaputra/paw-aparts/collision"
	"github.com/aclaputra/paw-aparts/components"
	"github.com/aclaputra/paw-aparts/config"
	"github.com/aclaputra/paw-aparts/tags"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdatePlayer(ecs *ecs.ECS) {
	playerEntry, _ := components.Player.First(ecs.World)
	player := components.Player.Get(playerEntry)
	playerObject := collision.GetObject(playerEntry)

	friction := 0.5
	accel := 0.5 + friction
	maxSpeed := 2.0
	jumpSpd := 10.0
	gravity := 0.75

	player.SpeedY += gravity
	if player.WallSliding != nil && player.SpeedY > 1 {
		player.SpeedY = 1
	}

	if player.WallSliding == nil {
		var isPlayerWalk bool
		if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxis(0, 0) > 0.1 {
			player.SpeedX += accel
			player.FacingRight = true
			isPlayerWalk = true
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxis(0, 0) < -0.1 {
			player.SpeedX -= accel
			player.FacingRight = false
			isPlayerWalk = true
		}

		if isPlayerWalk {
			player.Animation.State = "run"
		} else {
			player.Animation.State = "idle"
		}
	}

	if player.OnGround == nil {
		player.Animation.State = "running_jump"
	}

	if player.SpeedX > friction {
		player.SpeedX -= friction
	} else if player.SpeedX < -friction {
		player.SpeedX += friction
	} else {
		player.SpeedX = 0
	}

	if player.SpeedX > maxSpeed {
		player.SpeedX = maxSpeed
	} else if player.SpeedX < -maxSpeed {
		player.SpeedX = -maxSpeed
	}

	// Check for jumping.
	if inpututil.IsKeyJustPressed(ebiten.KeyX) || ebiten.IsGamepadButtonPressed(0, ebiten.GamepadButton0) || ebiten.IsGamepadButtonPressed(1, ebiten.GamepadButton0) {

		if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.GamepadAxis(0, 1) > 0.1 || ebiten.GamepadAxis(1, 1) > 0.1) && player.OnGround != nil && player.OnGround.HasTags("platform") {

			player.IgnorePlatform = player.OnGround

		} else {

			if player.OnGround != nil {
				player.SpeedY = -jumpSpd
			} else if player.WallSliding != nil {
				// WALLJUMPING
				player.SpeedY = -jumpSpd

				if player.WallSliding.Position.X > playerObject.Position.X {
					player.SpeedX = -4
				} else {
					player.SpeedX = 4
				}

				player.WallSliding = nil

			}

		}

	}
	dx := player.SpeedX

	if check := playerObject.Check(player.SpeedX, 0, "solid"); check != nil {
		dx = check.ContactWithCell(check.Cells[0]).X
		player.SpeedX = 0

		// If you're in the air, then colliding with a wall object makes you start wall sliding.
		if player.OnGround == nil {
			player.WallSliding = check.Objects[0]
		}

	}

	playerObject.Position.X += dx

	player.OnGround = nil

	dy := player.SpeedY

	dy = math.Max(math.Min(dy, 16), -16)
	checkDistance := dy
	if dy >= 0 {
		checkDistance++
	}

	if check := playerObject.Check(0, checkDistance, "solid", "platform"); check != nil {
		if platforms := check.ObjectsByTags("platform"); len(platforms) > 0 {

			platform := platforms[0]

			if platform != player.IgnorePlatform && player.SpeedY >= 0 && playerObject.Bottom() < platform.Position.Y+4 {
				dy = check.ContactWithObject(platform).Y
				player.OnGround = platform
				player.SpeedY = 0
			}

		}

		if solids := check.ObjectsByTags("solid"); len(solids) > 0 && (player.OnGround == nil || player.OnGround.Position.Y >= solids[0].Position.Y) {
			dy = check.ContactWithObject(solids[0]).Y
			player.SpeedY = 0

			if solids[0].Position.Y > playerObject.Position.Y {
				player.OnGround = solids[0]
			}

		}

		if player.OnGround != nil {
			player.WallSliding = nil
			player.IgnorePlatform = nil
		}

	}

	playerObject.Position.Y += dy

	wallNext := 1.0
	if !player.FacingRight {
		wallNext = -1
	}

	if c := playerObject.Check(wallNext, 0, "solid"); player.WallSliding != nil && c == nil {
		player.WallSliding = nil
	}

	player.HandleAnimationCount()
}

func DrawPlayer(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Player.Each(ecs.World, func(e *donburi.Entry) {
		player := components.Player.Get(e)
		o := collision.GetObject(e)
		if player.RespawnCooldown > 0 {
			player.RespawnCooldown--
		}

		if player.WallSliding != nil {
			player.Animation.State = "wall_sliding"
		}

		var (
			currentPlayerImg     *ebiten.Image
			aniTotalX, aniTotalY int
		)

		convertCatIdleAsset := func() {
			currentPlayerImg = assets.GetEbitenImage(assets.IdleCat_Png)
			aniTotalX, aniTotalY, player.Animation.FrameCount = 8, 1, 8
		}
		convertCatJumpAsset := func() {
			currentPlayerImg = assets.GetEbitenImage(assets.RunJumpCat_Png)
			aniTotalX, aniTotalY, player.Animation.FrameCount = 3, 1, 3
		}
		convertCatRunAsset := func() {
			currentPlayerImg = assets.GetEbitenImage(assets.RunCat_Png)
			aniTotalX, aniTotalY, player.Animation.FrameCount = 8, 1, 8 // know it by enum (should create enums or list of player anim states)
		}

		switch player.Animation.State {
		case "idle":
			convertCatIdleAsset()
		case "run":
			convertCatRunAsset()
		case "running_jump":
			convertCatJumpAsset()
		default:
			convertCatIdleAsset()
		}

		if player.SpeedY > 70 && player.RespawnCooldown == 0 {
			o.Position.X, o.Position.Y = (float64(config.C.Width)/2)-(80/2), float64(config.C.Height)-(128*4)
			player.RespawnCooldown = 60 * 3
		}

		op := ebiten.DrawImageOptions{}
		scale := 1.0
		var frameOX, frameOY int
		frameWidth, frameHeight :=
			currentPlayerImg.Bounds().Dx()/aniTotalX,
			currentPlayerImg.Bounds().Dy()/aniTotalY

		i := (player.GetAnimationCount() / (aniTotalX * aniTotalY)) % player.Animation.FrameCount

		var sx, sy int
		sx = frameOX + i*frameWidth
		sy = frameOY

		if player.FacingRight {
			op.GeoM.Scale(-scale, scale)
			op.GeoM.Translate(
				o.Position.X+float64(frameWidth)*scale,
				o.Position.Y+16,
			)
		} else {
			op.GeoM.Scale(scale, scale)
			op.GeoM.Translate(
				o.Position.X,
				o.Position.Y+16,
			)
		}

		screen.DrawImage(
			currentPlayerImg.SubImage(
				image.Rect(
					sx, sy, sx+frameWidth, sy+frameHeight,
				),
			).(*ebiten.Image), &op)
	})
}
