package config

type Config struct {
	Width     int
	Height    int
	GameTitle string
}

var C *Config

func init() {
	C = &Config{
		Width:     640,
		Height:    480,
		GameTitle: "Paw Aparts",
	}
}
