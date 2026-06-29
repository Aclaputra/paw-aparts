package assets

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed splash_logo.png
	SplashLogo_Png []byte
	//go:embed background.png
	Background_Png []byte
	//go:embed bg_game.png
	BGGame_Png []byte
	//go:embed bg_game_dark.png
	BGGameDark_Png []byte
	//go:embed fonts/excel.ttf
	ExcelFont []byte
	//go:embed cat/IDLE.png
	IdleCat_Png []byte
	//go:embed cat/JUMP.png
	JumpCat_Png []byte
	//go:embed cat/WALK.png
	WalkCat_Png []byte
	//go:embed cat/RUN.png
	RunCat_Png []byte
	//go:embed cat/RUNNING_JUMP.png
	RunJumpCat_Png []byte
	//go:embed music/game.ogg
	MusicGame []byte
	//go:embed music/main_menu.ogg
	MusicMainMenu []byte
	//go:embed music/result_menu.ogg
	MusicResultMenu []byte
)

// GetEbitenImage converts embedded []byte into *ebiten.Image
func GetEbitenImage(raw []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(raw))
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}
	return ebiten.NewImageFromImage(img)
}
